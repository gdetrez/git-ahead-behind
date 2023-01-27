Setup repo

  $ git clone --quiet $TESTDIR/testdata/test.git && cd test
  $ git checkout --quiet feature
  $ git checkout --quiet other-feature
  $ git checkout --quiet main

Run our tool

  $ export PATH=$TESTDIR:$PATH
  $ git ahead-behind
    feature        152dc1f8   11\xe2\x95\xba\xe2\x94\xbf\xe2\x94\x81101 (esc)
  * main           4520d372    0 \xe2\x94\x82 0   (esc)
    other-feature  293065ec    0 \xe2\x94\x9d 1   (esc)
