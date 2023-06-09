#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

bold=$(tput bold) || true
norm=$(tput sgr0) || true
red=$(tput setaf 1) || true
green=$(tput setaf 2) || true

fail() {
	echo "${red}$*${norm}."
	exit 1
}

echo "${DIR}"

echo "${bold}Running all tests...${norm}"
for testdir in "${DIR}"/*; do
	echo "${testdir}"
	if [[ -x "${testdir}/test.sh" ]]; then
		testname=$(basename "$testdir")
		echo "${bold}Running \"$testname\" test...${norm}"
		if cd "${testdir}"; then 
			if ./test.sh; then
				echo "${green}\"$testname\" test succeeded${norm}"
			else
				echo "${red}\"$testname\" test failed${norm}"
				FAILED=true
			fi
			cd ..
		else 
			echo "${red}\"$testname\" test failed${norm}"
				FAILED=true
		fi
	fi
done

if [ -n "${FAILED}" ]; then
	fail "There were test failures"
fi
echo "${green}Done. All test passed!${norm}"
exit 0