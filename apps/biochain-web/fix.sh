#!/bin/bash

# Directory where the script will start looking for index.js files
dir="presets/wind"

# Loop over all index.js files in the directory and its subdirectories
find "$dir" -name 'index.js' | while read -r file; do
    echo "Processing $file"
    # Use sed to perform the replacement
    # sed -i.bak '
    #     1s/^export default/export const ui =/
    #     $a\
    #     export default ui;
    # ' "$file"

    # Remove the backup file created by sed
    rm "${file}.bak"
done
