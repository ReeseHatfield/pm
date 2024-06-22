#!/bin/bash

# Find USB device details
lsusb_output=$(lsusb)
echo "$lsusb_output"

# Ask the user to input the Bus and Device numbers
echo "USB BUS NUM:"
read bus
echo "USB DEV NUM:"
read device

# Use udevadm to get the serial number
serial=$(udevadm info --name=/dev/bus/usb/$bus/$device) #  | grep SERIAL # awk -F'=' '{print $2}')

echo "Serial Number: $serial"
