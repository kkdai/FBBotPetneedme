PetNeedMe: Package to get Pet Adoption OpenData from Taiwan
==============

[![Join the chat at https://gitter.im/kkdai/PetNeedMe](https://badges.gitter.im/kkdai/PetNeedMe.svg)](https://gitter.im/kkdai/PetNeedMe?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

 [![GoDoc](https://godoc.org/github.com/kkdai/PetNeedMe?status.svg)](https://godoc.org/github.com/kkdai/PetNeedMe)  [![Build Status](https://travis-ci.org/kkdai/PetNeedMe.svg?branch=master)](https://travis-ci.org/kkdai/PetNeedMe)

A Package to get all Animal Adoption Open Data in every county in Taiwan.


A Console program to access Pet Information
=============

`go get github.com/kkdai/pnm/cmd/showpetcli`

Just run it `show[etcli`

- `n` for next Pet
- `p` for previous Pet
- `c` for next Cat
- `d` for next Dog


![](images/showpetcli.png)

Progress
=============

- Taipei (done)
- Other Counties (TBD)


Data Source:
=============

[「臺北市開放認養動物」API存取](http://data.taipei/opendata/datalist/datasetMeta/outboundDesc?id=6a3e862a-e1cb-4e44-b989-d35609559463&rid=f4a75ba9-7721-4363-884d-c3820b0b917c)

Contribution and Issue
=============

歡迎到 [issue](https://github.com/kkdai/LineBotTaipeiPets/issues) 寫下你的意見，或是一起來幫助我．



Inspired By
=============

- [Golang (heroku) で LINE Bot 作ってみる](http://qiita.com/dongri/items/ba150f04a98e96b160e7)
- [LINE BOT をとりあえずタダで Heroku で動かす](http://qiita.com/yuya_takeyama/items/0660a59d13e2cd0b2516)
- [petneed.me](https://github.com/jsleetw/petneed.me)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


致謝
---------------

感謝[g0v](http://g0v.tw/)的許多人不斷地提起這個專案，讓我可以注意到並且能夠一起幫忙．

- [Hackpad: Petneedme](https://g0v.hackpad.com/ep/pad/static/GOdHRgQpZSL)
- [Hackpad: 公立動物收容所資訊統整系統](https://g0v.hackpad.com/ep/pad/static/JBhVDOPxhxe) ([active one](https://g0v.hackpad.com/JBhVDOPxhxe))


License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

