#!/bin/bash
MEGAM_LOG="/var/log/megam/megamcib/opennebula.log"
echo "Step 1: you are running the test file. megam install is sleeping..." >> $MEGAM_LOG
sleep 10
exit 1
