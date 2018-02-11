#! /bin/sh
echo 'Ensure dependencies...'
cd src/tracer && dep ensure && cd ../..
echo 'Dependencies ensured.'
