/*-------------------------------------------------------------------------
 * This source code is confidential proprietary information which is
 * copyright (c) 2018, 2019 TrueBlocks, LLC (http://trueblocks.io)
 * All Rights Reserved
 *------------------------------------------------------------------------*/
#include "etherlib.h"
#include "options.h"

//-----------------------------------------------------------------------
int main(int argc, const char* argv[]) {
    nodeNotRequired();
    acctlib_init(defaultQuitHandler);

    COptions options;
    if (!options.prepareArguments(argc, argv))
        return 0;

    for (auto command : options.commandLines) {
        if (!options.parseArguments(command))
            continue;

        // Clean the staging folder since we want to start with clean scrapes for both
        // never scraped or unfinished (partial) scrapes
        cleanFolder(getMonitorPath("", FM_STAGING));

        if (options.visitTypes & VIS_FINAL)
            forEveryFileInFolder(indexFolder_finalized, visitFinalIndexFiles, &options);

        if (options.visitTypes & VIS_STAGING)
            forEveryFileInFolder(indexFolder_staging, visitStagingIndexFiles, &options);

        if (options.visitTypes & VIS_UNRIPE)
            forEveryFileInFolder(indexFolder_unripe, visitUnripeIndexFiles, &options);

        options.moveToProduction();
    }

    acctlib_cleanup();
    return 0;
}

//--------------------------------------------------------------------------------
void doMoveFile(const string_q& from, const string_q& to) {
    if (verbose)
        cerr << "Moving " << cTeal << from << cOff << " to " << cTeal << to << cOff << endl;
    if (fileExists(from))
        moveFile(from, to);
}

//--------------------------------------------------------------------------------
void COptions::moveToProduction(void) {
    for (auto acct : monitors) {
        if (acct.fm_mode == FM_STAGING) {
            acct.fm_mode = FM_PRODUCTION;
            if (acct.tx_cache) {
                acct.tx_cache->Release();
                delete acct.tx_cache;
                acct.tx_cache = NULL;
            }
            lockSection(true);
            doMoveFile(getMonitorPath(acct.address, FM_STAGING), getMonitorPath(acct.address));
            doMoveFile(getMonitorLast(acct.address, FM_STAGING), getMonitorLast(acct.address));
            doMoveFile(getMonitorExpt(acct.address, FM_STAGING), getMonitorExpt(acct.address));
            doMoveFile(getMonitorBals(acct.address, FM_STAGING), getMonitorBals(acct.address));
            lockSection(false);
        }
    }
}
