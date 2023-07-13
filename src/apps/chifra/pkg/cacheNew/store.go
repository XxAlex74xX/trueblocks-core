package cacheNew

import (
	"bytes"
	"errors"
	"os"
	"path"
	"sync"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cacheNew/locations"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// In verbose mode we print cache errors. It's useful for debugging.
var verboseMode = false

func init() {
	if os.Getenv("CACHE_VERBOSE") == "true" {
		verboseMode = true
	}
}

var defaultStore *Store
var defaultStoreOnce sync.Once

type Store struct {
	resolvedPaths map[Locator]string
	location      Storer
	rootDir       string
}

func NewStore(options *StoreOptions) (*Store, error) {
	location, err := options.location()
	if err != nil {
		return nil, err
	}
	return &Store{
		location: location,
		rootDir:  options.rootDir(),
	}, nil
}

func DefaultStore() (*Store, error) {
	var err error
	defaultStoreOnce.Do(func() {
		defaultStore, err = NewStore(nil)
	})
	return defaultStore, err
}

func (s *Store) resolvePath(value Locator) (resolved string, err error) {
	if cachedPath, ok := s.resolvedPaths[value]; ok {
		return cachedPath, nil
	}

	id := value.CacheId()
	directory, extension := value.CacheLocation()
	if directory == "" || extension == "" {
		err = errors.New("empty CacheLocation")
		return
	}
	resolved = path.Join(s.rootDir, directory, (id + "." + extension))
	return
}

// Write options that we might need in the future
type WriteOptions interface{}

// Write saves value to a location defined by options.Location. If options is nil,
// then FileSystem is used. The value has to implement Locator interface, which
// provides information about in-cache path and ID.
func (s *Store) Write(value Locator, options *WriteOptions) (err error) {
	itemPath, err := s.resolvePath(value)
	if err != nil {
		printErr("write resolving path", err)
		return
	}

	writer, err := s.location.Writer(itemPath)
	if err != nil {
		printErr("getting writer", err)
		return
	}
	defer writer.Close()

	buffer := new(bytes.Buffer)
	item := NewItem(buffer)
	if err = item.Encode(value); err != nil {
		printErr("encoding", err)
		return
	}
	_, err = buffer.WriteTo(writer)

	return
}

// Read options that we might need in the future
type ReadOptions interface{}

// Read retrieves value from a location defined by options.Location. If options is nil,
// then FileSystem is used. The value has to implement Locator interface, which
// provides information about in-cache path
func (s *Store) Read(value Locator, options *ReadOptions) (err error) {
	itemPath, err := s.resolvePath(value)
	if err != nil {
		printErr("read resolving path", err)
		return
	}

	reader, err := s.location.Reader(itemPath)
	if err != nil {
		printErr("getting reader", err)
		return
	}
	defer reader.Close()

	buffer := new(bytes.Buffer)
	if _, err = buffer.ReadFrom(reader); err != nil {
		printErr("reading", err)
		return
	}

	item := NewItem(buffer)
	err = item.Decode(value)
	if err != nil {
		printErr("decoding", err)
	}
	return
}

func (s *Store) Stat(value Locator) (result *locations.ItemInfo, err error) {
	itemPath, err := s.resolvePath(value)
	if err != nil {
		printErr("stat resolving path", err)
		return
	}
	result, err = s.location.Stat(itemPath)
	if err != nil {
		printErr("stat", err)
	}
	return
}

func (s *Store) Remove(value Locator) error {
	itemPath, err := s.resolvePath(value)
	if err != nil {
		printErr("remove resolving path", err)
		return err
	}
	err = s.location.Remove(itemPath)
	if err != nil {
		printErr("removing", err)
	}
	return err
}

type DecacheFunc func(*locations.ItemInfo) bool

func (s *Store) Decache(locators []Locator, processor DecacheFunc) (err error) {
	for _, locator := range locators {
		// directory, extension := locator.CacheLocation()
		// cachePath := path.Join(directory, locator.CacheId(), extension)
		stats, err := s.Stat(locator)
		if err != nil {
			return err
		}
		// If processor returns false, we don't want to remove this item from the cache
		if !processor(stats) {
			continue
		}
		if err := s.Remove(locator); err != nil {
			return err
		}
	}

	return nil
}

func printErr(desc string, err error) {
	if !verboseMode {
		return
	}

	logger.Warn("cache error:", desc, ":", err)
}
