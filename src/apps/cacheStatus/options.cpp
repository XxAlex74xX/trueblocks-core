/*-------------------------------------------------------------------------
 * This source code is confidential proprietary information which is
 * copyright (c) 2018, 2019 TrueBlocks, LLC (http://trueblocks.io)
 * All Rights Reserved
 *------------------------------------------------------------------------*/
#include "options.h"

//---------------------------------------------------------------------------------------------------
static const COption params[] = {
// BEG_CODE_OPTIONS
    COption("modes", "", "list<enum[index|monitors|names|abis|blocks|transactions|traces|slurps|prices|some*|all]>", OPT_POSITIONAL, "which data to retrieve"),
    COption("details", "d", "", OPT_SWITCH, "include details about items found in monitors, slurps, abis, or price caches"),
    COption("list", "l", "", OPT_SWITCH, "display results in Linux ls -l format (assumes --detail)"),
    COption("report", "r", "", OPT_SWITCH, "show a summary of the current status of the blockchain and TrueBlocks scrapers"),
    COption("get_config", "g", "", OPT_HIDDEN | OPT_SWITCH, "returns JSON data of the editable configuration file items"),
    COption("set_config", "s", "", OPT_HIDDEN | OPT_SWITCH, "accepts JSON in an env variable and writes it to configuration files"),
    COption("start", "S", "<blknum>", OPT_HIDDEN | OPT_FLAG, "first block to process (inclusive)"),
    COption("end", "E", "<blknum>", OPT_HIDDEN | OPT_FLAG, "last block to process (inclusive)"),
    COption("", "", "", OPT_DESCRIPTION, "Report on status of one or more TrueBlocks caches."),
// END_CODE_OPTIONS
};
static const size_t nParams = sizeof(params) / sizeof(COption);

//---------------------------------------------------------------------------------------------------
bool COptions::parseArguments(string_q& command) {
    ENTER4("parseArguments");
    if (!standardOptions(command))
        EXIT_NOMSG(false);

// BEG_CODE_LOCAL_INIT
    bool report = false;
    bool get_config = false;
    bool set_config = false;
// END_CODE_LOCAL_INIT

    blknum_t latest = getLastBlock_client();

    Init();
    explode(arguments, command, ' ');
    for (auto arg : arguments) {
        if (false) {
            // do nothing -- make auto code generation easier
// BEG_CODE_AUTO
        } else if (arg == "-d" || arg == "--details") {
            details = true;

        } else if (arg == "-l" || arg == "--list") {
            list = true;

        } else if (arg == "-r" || arg == "--report") {
            report = true;

        } else if (arg == "-g" || arg == "--get_config") {
            get_config = true;

        } else if (arg == "-s" || arg == "--set_config") {
            set_config = true;

        } else if (startsWith(arg, "-S:") || startsWith(arg, "--start:")) {
            if (!confirmBlockNum("start", start, arg, latest))
                return false;

        } else if (startsWith(arg, "-E:") || startsWith(arg, "--end:")) {
            if (!confirmBlockNum("end", end, arg, latest))
                return false;

        } else if (startsWith(arg, '-')) {  // do not collapse

            if (!builtInCmd(arg)) {
                return usage("Invalid option: " + arg);
            }

// END_CODE_AUTO
        } else {
            string_q permitted = params[0].description;
            replaceAny(permitted, "[]*", "|");
            if (!contains(permitted, "|" + arg + "|"))
                return usage("Provided value for 'mode' (" + arg + ") not " + substitute(params[0].description, "enum", "") + ". Quitting.");
            mode += (arg + "|");
        }
    }

    if (start == NOPOS)
        start = 0;

    if (get_config && set_config)
        return usage("Please chose only one of --set_config and --get_config. Quitting...");

    if (set_config) {
        mode = "set";
        isConfig = true;
    } else if (get_config) {
        isConfig = true;
    }

    if (!isConfig) {
        if (mode.empty() || contains(mode, "some"))
            mode = "index|monitors|names|slurps|prices";

        if (contains(mode, "all"))
            mode = "index|monitors|names|abis|blocks|transactions|traces|slurps|prices";

        mode = "|" + trim(mode, '|') + "|";
    }

    if (!details) {
        HIDE_FIELD(CMonitorCache, "items");
        HIDE_FIELD(CSlurpCache, "items");
        HIDE_FIELD(CPriceCache, "items");
        HIDE_FIELD(CAbiCache, "items");
    } else {
        loadHashes(indexHashes, "finalized");
        loadHashes(bloomHashes, "blooms");
    }

    if (report) {
        if (exportFmt == TXT1 || exportFmt == CSV1) {
            cout << scraperStatus(false);
            return false;
        }
    }

    EXIT_NOMSG(true);
}

//---------------------------------------------------------------------------------------------------
void COptions::Init(void) {
    registerOptions(nParams, params);
    optionOn(OPT_PREFUND | OPT_OUTPUT);

// BEG_CODE_INIT
    details = false;
    list = false;
    start = NOPOS;
    end = NOPOS;
// END_CODE_INIT

    isConfig = false;
    mode = "";

    char hostname[HOST_NAME_MAX];  gethostname(hostname, HOST_NAME_MAX);
    char username[LOGIN_NAME_MAX]; getlogin_r(username, LOGIN_NAME_MAX);
    if (!getEnvStr("DOCKER_MODE").empty()) {
        memset(username, 0, LOGIN_NAME_MAX);
        strncpy(username, "nobody", 6);
    }

    if (isTestMode()) {
        status.host = "--hostname-- (--username--)";
        status.rpc_provider = status.api_provider = status.balance_provider = "--providers--";
    } else {
        status.host = string_q(hostname) + " (" + username + ")";
        status.rpc_provider = getGlobalConfig()->getConfigStr("settings", "rpcProvider", "http://localhost:8545");
        status.api_provider = getGlobalConfig()->getConfigStr("settings", "apiProvider", "http://localhost:8080");
        status.balance_provider = getGlobalConfig()->getConfigStr("settings", "balanceProvider", status.api_provider);
    }
    if (!isNodeRunning()) {
        status.client_version = "Not running";
    } else {
        status.client_version = (isTestMode() ? "Parity version" : getVersionFromClient());
    }
    status.trueblocks_version = getVersionStr();
    status.is_scraping = contains(listRunning("chifra scrape"), "chifra scrape");
    if (isTestMode())
        status.is_scraping = false;
}

//---------------------------------------------------------------------------------------------------
COptions::COptions(void) {
    setSorts(GETRUNTIME_CLASS(CBlock), GETRUNTIME_CLASS(CTransaction), GETRUNTIME_CLASS(CReceipt));
    Init();

    CStatus::registerClass();
    CCache::registerClass();
    CIndexCache::registerClass();
    CIndexCacheItem::registerClass();
    CAbiCache::registerClass();
    CChainCache::registerClass();
    CMonitorCache::registerClass();
    CMonitorCacheItem::registerClass();
    CSlurpCache::registerClass();
    CPriceCache::registerClass();
    CPriceCacheItem::registerClass();
    CAbiCacheItem::registerClass();
    CConfiguration::registerClass();
    CConfigFile::registerClass();
    CConfigGroup::registerClass();
    CConfigItem::registerClass();

    HIDE_FIELD(CAccountWatch,  "statement");
    HIDE_FIELD(CAccountWatch,  "stateHistory");
    HIDE_FIELD(CAccountWatch,  "curBalance");
    HIDE_FIELD(CAccountWatch,  "abi_spec");

    HIDE_FIELD(CAccountName,   "is_contract");
    HIDE_FIELD(CAccountName,   "is_private");
    HIDE_FIELD(CAccountName,   "is_shared");
    UNHIDE_FIELD(CAccountName, "nRecords");
    UNHIDE_FIELD(CAccountName, "nAppearances");
    UNHIDE_FIELD(CAccountName, "sizeInBytes");
    UNHIDE_FIELD(CAccountName, "firstAppearance");
    UNHIDE_FIELD(CAccountName, "latestAppearance");

    minArgs = 0;
}

//--------------------------------------------------------------------------------
COptions::~COptions(void) {
}

//--------------------------------------------------------------------------------
string_q COptions::postProcess(const string_q& which, const string_q& str) const {
    if (which == "options") {
        return substitute(str, "modes", "<mode> [mode...]");

    } else if (which == "notes" && (verbose || COptions::isReadme)) {
        // do nothing

    }
    return str;
}

//--------------------------------------------------------------------------------
void loadHashes(CIndexHashMap& map, const string_q& which) {
    string_q hashFn = configPath("ipfs-hashes/" + which + ".txt");
    if (!fileExists(hashFn)) {
        cerr << "Hash file (" << hashFn << ") not found." << endl;
    } else {
        string_q contents = asciiFileToString(hashFn);
        CStringArray lines;
        explode(lines, contents, '\n');
        for (auto line : lines) {
            line = substitute(substitute(line, ".bin", ""), ".bloom", "");
            CStringArray parts;
            explode(parts, line, ' ');
            map[parts[2]] = parts[1];
        }
    }
}