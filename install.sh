#!/bin/bash

echo "Installing squigglelog..."

echo "Linking entries folder to $HOME/~entries"
if [ ! -e $HOME/~entries ]; then
	ln -s $PWD/entries/ $HOME/~entries
else
	echo "Folder already exists, skipping..."
fi

echo "Installing utility scripts..."
if [ ! -e $HOME/bin ]; then
	EM="\033[1;33m"
	NC="\033[0m"
	echo "Creating bin/ folder"
	mkdir $HOME/bin
	echo -e "${EM}You may have to add this to your ~/.bashrc file:"
	echo -e "export PATH=${PATH}:$HOME/bin$NC"
fi
cp util/* $HOME/bin/

echo "Building squigglelog..."
go build squigglelog.go

echo "Done!"
echo
echo "To get started, type squiggle.sh and write your first entry."
