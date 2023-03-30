#!/bin/sh

while read p || [ -n "$p" ] 
do  
sed -i '' "/${p//\//\\/}/d" ./coverage.txt
done < ./.coverage_ignore