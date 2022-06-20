#!/bin/bash
#
# based on https://stackoverflow.com/questions/40450238/parse-a-changelog-and-extract-changes-for-a-version
#
# works with changelogs in the format of https://keepachangelog.com
#
# usage: ./extractVersionChanges.sh "1.0.1" ./CHANGELOG.md ./CHANGELOG-CHANGES.md
#
awk -v ver="$1" '
        /^#+ \[/ { if (p) { exit }; if ($2 == "["ver"]") { p=1; next} } p && NF
    ' "$2" >"$3"
