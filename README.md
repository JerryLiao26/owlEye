# owlEye
[![Build Status](https://travis-ci.org/JerryLiao26/owlEye.svg?branch=master)](https://travis-ci.org/JerryLiao26/owlEye)
[![Go Report Card](https://goreportcard.com/badge/github.com/JerryLiao26/owlEye)](https://goreportcard.com/report/github.com/JerryLiao26/owlEye)
[![codebeat badge](https://codebeat.co/badges/dda4dfb4-d384-40c6-9168-8fe6edde8945)](https://codebeat.co/projects/github-com-jerryliao26-owleye-master)
[![License](https://img.shields.io/github/license/JerryLiao26/owlEye.svg)](https://opensource.org/licenses/MIT)

Enter a piece of description, and ```owlEye``` will pick the target image for you.

## Requirements
- ```owlEye``` uses [lorca](https://github.com/zserge/lorca) as HTML UI wrapper, therefore you will need Chrome/Chromium >= 70 installed on your system.
- ```owlEye``` gets user's home directory with ```os.UserHomeDir()``` introduced in Go 1.12 so the Go version is also limited.

## External Servers
For convenience of using Machine Learning models, ```owlEye``` does not build internal processing servers. You have to run compatible ones yourself then configure ```owlEye``` to run with them. Demos are listed below:
- [owlEye-text-server](https://github.com/JerryLiao26/owlEye-text-server)