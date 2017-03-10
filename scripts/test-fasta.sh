#!/bin/bash

##
##  Live tests of the binary, primarily to test flag usage and error production
##  Those are difficult to test internally, so this is a final set of tests.
##  Changes in the interface may induce failures, particularly around the "usage"
##  statement.
##

TMP=$(mktemp -d ./testdata/livetest.XXXXXX)

function testfail {
  echo "$@"
  # clean up.  Remove/comment this to debug failing tests
  rm -r $TMP
  exit 1
}

t=1
# Test 1, successful assembly, no flags
./fasta-example testdata/match-success > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  diff testdata/live/correct $TMP/$t.out || testfail "Test $t: output match failure"
  diff /dev/null $TMP/$t.error || testfail "Test $t: produced stderr when none was expected"
  echo "Live test $t complete: OK"
else
  testfail "Test $t: expected success, got error $(cat $TMP/$t.error)"
fi

t=2
# Test 2, successful assembly, fasta format, wrap 4
./fasta-example -t test -w 4 testdata/match-success > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  diff testdata/live/fasta $TMP/$t.out || testfail "Test $t: output match failure"
  diff /dev/null $TMP/$t.error || testfail "Test $t: produced stderr when none was expected"
  echo "Live test $t complete: OK"
else
  testfail "Test $t: expected success, got error $(cat $TMP/$t.error)"
fi

t=3
# Test 3, successful assembly, debug output
./fasta-example -d testdata/match-success  > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  diff testdata/live/correct $TMP/$t.out || testfail "Test $t: output match failure"
  if ! [ -s $TMP/$t.error ] ; then
    testfail "Test $t: produced no stderr with debug enabled"
  fi
  echo "Live test $t complete: OK"
else
  testfail "Test $t: expected success, got error $(cat $TMP/$t.error)"
fi


# Failure tests
t=4
# Test 4, no filename fail/usage
./fasta-example > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  testfail "Test $t: failed to produce error"
else
  diff testdata/live/usage $TMP/$t.error || testfail "Test $t: usage string match failure"
  echo "Live test $t complete: OK"
fi

t=5
# Test 5, all flags with -h, produce usage
./fasta-example -h -d -w 4 -t stuff testdata/match-success> $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  testfail "Test $t: failed to produce error"
else
  diff testdata/live/usage $TMP/$t.error || testfail "Test $t: usage string match failure"
  echo "Live test $t complete: OK"
fi

t=6
# Test 6, two filenames, fail/usage
./fasta-example file1 file2 > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  testfail "Test $t: failed to produce error"
else
  diff testdata/live/usage $TMP/$t.error || testfail "Test $t: usage string match failure"
  echo "Live test $t complete: OK"
fi

t=7
# Test 7, bad data, fail
./fasta-example testdata/2dirty.txt > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  testfail "Test $t: failed to produce error"
else
  echo "Live test $t complete: OK"
fi

t=8
# Test 8, unmatchable, fail
./fasta-example testdata/match-fail > $TMP/$t.out 2> $TMP/$t.error
if  [ $? -eq 0 ] ; then
  testfail "Test $t: failed to produce error"
else
  echo "Live test $t complete: OK"
fi

rm -r $TMP
