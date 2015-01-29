#!/bin/bash

entriesDir=$HOME/~entries/
vi "$entriesDir$(date +'%Y%m%d')" "+set wrap" "+set linebreak"
echo "Done!"
