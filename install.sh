#!/bin/bash
# Use this file to install dependencies as well
# as the application into terminal

# Performing go mod tidy to install unavailable dependencies
printf "\nDownloading unavailable go packages\n"
if go mod tidy; then 
    echo "Installed go packages"
else
    echo "Could not install go packages"
    exit 1
fi

# Doing a go build to generate binary
printf "\nGenerating binary\n"
if go build; then 
    echo "Built go binary - calc"
else
    echo "Could not build go binary"
    exit 1
fi

# Installing new binary with go install
printf "\nInstalling calc binary\n"
if go install calc; then 
    echo "Installed calc binary"
else
    echo "Could not install calc binary"
    exit 1
fi

# Checking installation
printf "\nVersion check\n"
if calc version; then 
    echo
else
    echo "Could not access calc"
    exit 1
fi

calc