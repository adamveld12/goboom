# Goboom

[![GoDoc](https://godoc.org/github.com/adamveld12/goboom?status.svg)](http://godoc.org/github.com/adamveld12/goboom)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/adamveld12/goboom)
[![Gocover](http://gocover.io/_badge/github.com/adamveld12/goboom)](http://gocover.io/github.com/adamveld12/goboom)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamveld12/goboom)](https://goreportcard.com/report/github.com/adamveld12/goboom)
[![Build Status](https://semaphoreci.com/api/v1/adamveld12/goboom/branches/master/badge.svg)](https://semaphoreci.com/adamveld12/goboom)


Boomerang beacon HTTP server.

You can write custom validators and exporters, allowing you to pipe the beacons into any kind of backend you wish.

## How to console

```sh
goboom -address "127.0.0.1:3000" -origin '.*\\.example\\.com' -url "/beacon"
```

## How to library

```golang
func main() {
	gb := goboom.Handler {
		Exporter: goboom.ConsoleExporter(os.Stdout),
	}

	log.Fatal(http.ListenAndServe("127.0.0.1:3000", gb))
}
```

## How to contribute

Start the server using `make run`.

Run the test watcher using `make dev`.

Make sure to run `make pc` before you push.

You can setup boomerang on a webpage and point boomerang at your local server, or you can visit 
[http://boomerang-test.surge.sh/](http://boomerang-test.surge.sh/) which will send an HTTP `POST` 
to `localhost:3000/beacon` making it easy for you to test things.

Or you can just test using curl commands like so:

```
curl --request POST \
  --url http://127.0.0.1:3000/beacon \
  --header 'content-type: application/x-www-form-urlencoded' \
  --header 'origin:  http://local.example.com:3000' \
  --header 'referer:  http://local.example.com:3000' \
  --header 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.79 Safari/537.36' \
  --data 'mob.etype=4g&mob.dl=3.15&mob.rtt=50&c.e=jj4jq4a4&c.tti.m=lt&rt.start=manual&rt.bmr=235%2C297&rt.blstart=1530552979318&pt.fp=1133&pt.fcp=1133&nt_nav_st=1530552978220&nt_fet_st=1530552978223&nt_dns_st=1530552978223&nt_dns_end=1530552978223&nt_con_st=1530552978223&nt_con_end=1530552978223&nt_req_st=1530552978226&nt_res_st=1530552978928&nt_res_end=1530552978940&nt_domloading=1530552978958&nt_domint=1530552981354&nt_domcontloaded_st=1530552981355&nt_domcontloaded_end=1530552981365&nt_domcomp=1530552983476&nt_load_st=1530552983476&nt_load_end=1530552983481&nt_unload_st=1530552978934&nt_unload_end=1530552978935&nt_enc_size=62880&nt_dec_size=311675&nt_trn_size=63313&nt_protocol=http%2F1.1&nt_first_paint=1530552979353&nt_red_cnt=0&nt_nav_type=1&restiming=%7B%22http%22%3A%7B%22%3A%2F%2F%22%3A%7B%22local%22%3A%7B%22.liveauctioneers.com%3A3000%2F%22%3A%226%2Ck1%2Cjp%2C6%2C3%2C%2C3%2C3%2C3*11cio%2Cc1%2C5byz%22%2C%22host%3A3001%2Fdist%2F%22%3A%7B%22main.546117f0e75a8cc708a0.js%22%3A%222l7%2Ckc*24*42%22%2C%22HomePage.546117f0e75a8cc708a0.js%22%3A%222l8%2Cdi*21*42%22%7D%7D%2C%22www.googletagmanager.com%2Fgtm.js%3Fid%3DGTM-MGNHR3N%22%3A%223ve%2C2r*21%22%7D%2C%22s%3A%2F%2F%22%3A%7B%22p1.liveauctioneers.com%2F%22%3A%7B%221537%2F123392%2F62778%22%3A%7B%223%22%3A%7B%2215_1_x.jpg%3Fversion%3D1528480713%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04p%2C3g%2Cfn%2C23m%2C6f%2C4q%7C1lb%2Ci4%2Cho%2Cdv*13we%2C5q%22%2C%2291_1_x.jpg%3Fversion%3D1530223057%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04q%2C23%2Ckd%2C24b%2Cap%2C4q%7C1lb%2Cic%2Ci4%2Cdv*11pu%2C5q%22%7D%2C%22299_1_x.jpg%3Fversion%3D1528480713%26width%3D340%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*09g%2C9d%2Cfn%2C1t6%2C9j%2C9g%7C1lb%2Chj%2Cg2%2Cc6*1ayw%2C86%22%7D%2C%222925%2F124249%2F63191%22%3A%7B%227%22%3A%7B%2226_1_x.jpg%3Fversion%3D1530257358%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04q%2C4q%2Ckd%2C1mv%7C1lb%2Cg7%2Cg2%2Cc5*12bo%2C6p%22%2C%2240_1_x.jpg%3Fversion%3D1530257358%26width%3D340%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*09g%2C9g%2Cfn%2C1d1%7C1lb%2Cfb%2Ce3%2Cay*18a6%2C60%22%7D%2C%22690_1_x.jpg%3Fversion%3D1530257358%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04p%2C4p%2Cfn%2C1mw%2C4q%2C4q%7C1lb%2Cfq%2Cfi%2Cb9*1284%2C5m%22%7D%2C%225584%2F123992%2F630768%22%3A%7B%2207_1_x.jpg%3Fversion%3D1530056507%26width%3D340%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*09g%2C9g%2Cfn%2Cgt%7C1lb%2Cbp%2Cau%2C9e*11po%2C8b%22%2C%2216_1_x.jpg%3Fversion%3D1530056507%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04q%2C4q%2Ckd%2Cqn%7C1lb%2Cc1%2Cbe%2C9f*1q8%2C5p%22%2C%2264_1_x.jpg%3Fversion%3D1530056507%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04p%2C4p%2Cfn%2Cqo%2C4q%2C4q%7C1lb%2Cbq%2Cb1%2C9e*114y%2C5s%22%7D%2C%22989%2F124264%2F63204%22%3A%7B%221%22%3A%7B%2279_1_x.jpg%3Fversion%3D1530208390%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04p%2C4p%2Cfn%2C16s%2C4q%2C4q%7C1lb%2Cbp%2Cb0%2C9g*114a%2C70%22%2C%2294_1_x.jpg%3Fversion%3D1530208390%26width%3D170%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*04q%2C4q%2Ckd%2C16r%7C1lb%2Ce4%2Cdp%2Cag*11g6%2C7a%22%7D%2C%22806_1_x.jpg%3Fversion%3D1530208390%26width%3D340%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22*09g%2C6o%2Cfn%2Cyb%2Cdd%2C9g%7C1lb%2Cdp%2Cbg%2C9g*1c3c%2C81%22%7D%2C%22dist%2F%22%3A%7B%22f%22%3A%7B%22onts%2Ffast-fonts%2F%22%3A%7B%223600b37f-2bf1-45f3-be3a-03365f16d9cb.woff2%22%3A%222l7%2C9x%2C20%2Ch*1r88%2C5d*42%22%2C%22febf3d0c-873f-4285-8ab4-77c31b26e747.woff2%22%3A%222l8%2C9y%2C3y%2Cg*1dus%2C51*42%22%2C%22b0868b4c-234e-47d3-bc59-41ab9de3c0db.woff2%22%3A%222l8%2C9x%2C3x%2Cg*1v5g%2C2o*42%22%7D%2C%22lags%2F%22%3A%7B%22US.png%22%3A%22*0m%2Co%2C1mf%2Cp0%2Cw%2Cw%7C1v0%2C8p%2C8n%2C5y*14c%2C7k%22%2C%22IT.png%22%3A%22*0m%2Co%2C1wr%2Cp0%2Cw%2Cw%7C1v0%2C1kx%2Cdz%2C61*13f%2C5o%22%2C%22NL.png%22%3A%22*0m%2Co%2C273%2C1br%2Cw%2Cw%7C1v1%2C1kx%2Cdz%2C6k*13i%2C7z%22%7D%7D%2C%22images%2F%22%3A%7B%22logo.svg%22%3A%22*0n%2C51%2C1p%2Cff%2C12%2C8c%7C1lb%2Cb6%2Ca3%2Cp*1vp%2C5w%2C13p%22%2C%22IPhone_Small.png%22%3A%22*02s%2C28%2C2u4%2Cj4%7C1v1%2C1kx%2Ce4%2C7z*116o%2C75%22%2C%22GooglePlay.png%3Fformat%3Dpjpg%26auto%3Dwebp%3Fformat%3Dpjpg%26auto%3Dwebp%22%3A%2242lf%2Cmf%2C1t%2C9*12fy%2C7o%22%2C%22DownloadAppStore.png%3Fformat%3Dpjpg%26auto%3Dwebp%3Fformat%3Dpjpg%26auto%3Dwebp%22%3A%2242lf%2Cme%2C1t%2C9*12wo%2C6v%22%2C%22desktop-%22%3A%7B%22a%22%3A%7B%22rt.jpg%3Fheight%3D382%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22132h%2C80%2C2k%2Cc*1apo%2C66%22%2C%22sian.jpg%3Fheight%3D382%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22132i%2C7y%2C2i%2Cc*1ag0%2C6b%22%7D%2C%22jewelry.jpg%3Fheight%3D382%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22132i%2Ccm%2C9f%2Cb*17m6%2C68%22%2C%22collectibles.jpg%3Fheight%3D382%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22132i%2C7x%2C1q%2Cc*159e%2C7i%22%2C%22memorabilia.jpg%3Fheight%3D382%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22132i%2Ccl%2C8m%2Cc*1432%2C7e%22%2C%22fashion.jpg%3Fheight%3D382%26format%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%22132j%2Ccj%2C8l%2Cb*17we%2C7t%22%7D%2C%22arrow-%22%3A%7B%22next.png%3Fauto%3Dwebp%22%3A%224350%2Cdb%2Cb0%2C6*1c8%2C6r%22%2C%22prev.png%3Fauto%3Dwebp%22%3A%22435e%2Cd1%2Cax%2C16*19s%2C6k%22%7D%2C%22hunt.jpg%3Fformat%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%224352%2Cm1%2Ccm%2C1i*1174k%2C8a%22%2C%22search.jpg%3Fformat%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%224354%2Cpb%2Cm0%2C6k*1v6u%2C7t%22%2C%22treasure.jpg%3Fformat%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%224356%2Cti%2Cp9%2C7e*1vzq%2C6i%22%2C%22explore.jpg%3Fformat%3Dpjpg%26auto%3Dwebp%26quality%3D50%22%3A%224358%2Cd5%2Cb2%2C1c*18tu%2C60%22%7D%7D%7D%2C%22www.liveauctioneers.com%2Fcodex.packages%2Fbiddable-widget%2F0.0.12.js%22%3A%223la%2Ca4*21%22%2C%22js.stripe.com%2Fv2%2F%22%3A%7B%22%7C%22%3A%2232ex%2Col%2C1m%2C4*1fw6%2Cae%2Cxe9*21%22%2C%22channel.html%3Fstripe_xdm_e%3Dhttp%253A%252F%252Flocal.liveauctioneers.com%253A3000%26stripe_xdm_c%3Ddefault943978%26stripe_xdm_p%3D1%23__stripe_transport__%22%3A%22*046%2C8c%2C-1jk%7Ca3do%2Cck%2C20%2Cf*1fk%2C72%2Cfd%22%2C%22m%2Fouter.html%23referrer%3D%26title%3DBid%2520in%2520Online%2520Auctions%2520-%2520LiveAuctioneers%26url%3Dhttp%253A%252F%252Flocal.liveauctioneers.com%253A3000%252F%26muid%3Da12e4095-3f12-43bf-91c2-08075f06179b%26sid%3Dca0dfce8-ee99-4ddd-ac00-ad2559d02a71%26preview%3Dfalse%26%22%3A%22*01%2C1%2C3l2%7Ca425%2C1t%2C1i%2C4*1as%2C50%2C96%22%7D%2C%22images.liveauctioneers.com%2Fstatic%2F%22%3A%7B%22sddefault.jpg%3Fwidth%3D560format%3Dpjpg%26auto%3Dwebp%22%3A%22434l%2Ccc%2Caa%2C3*1djk%2C6r%22%2C%22youtube-%22%3A%7B%22red.png%3Fwidth%3D72%26auto%3Dwebp%22%3A%22434l%2Ccm%2Cax%2C5*1ri%2C6q%22%2C%22gray.png%3Fwidth%3D72%26auto%3Dwebp%22%3A%22434l%2Ccn%2Cax%2C5*1wk%2C68%22%7D%7D%7D%7D%7D&spa.missed=1&rt.tstart=1530552978220&rt.bstart=1530552982463&rt.end=1530552983481&t_resp=709&t_page=4552&t_done=5261&t_other=boomerang%7C766%2Cboomr_fb%7C4243%2Cboomr_ld%7C1098%2Cboomr_lat%7C3145&u=http%3A%2F%2Flocal.liveauctioneers.com%3A3000%2F&http.initiator=spa_hard&r=http%3A%2F%2Flocal.liveauctioneers.com%3A3000%2F&v=1.0.0&vis.st=visible&ua.plt=Linux%20x86_64&ua.vnd=Google%20Inc.&pid=3cfp9u9i&if=&c.t.longtask=00001&c.t.fps=06457666767666&c.t.mem=1*a*0%2C5yc1s%2C5yc1s%2C5yc1s&c.t.domsz=1*a*0%2Cedkc%2Cedkc%2Cedkc&c.t.domln=1*a*0%2C1aj%2C1aj%2C1aj&c.tti.vr=3145&c.tti=5287&c.lt.n=1&c.lt.tt=51.40000000028522&c.lt=~(~(d~'\''1f~n~8~s~'\''420))&c.f=60&c.f.d=1292&c.f.m=4&c.f.s=jj4jq84o&dom.res=44&dom.doms=6&mem.total=10000000&mem.limit=2330000000&mem.used=10000000&scr.xy=2327x1309&scr.bpp=24%2F24&scr.orn=0%2Flandscape-primary&scr.dpx=1.649999976158142&cpu.cnc=8&bat.lvl=1&dom.ln=1675&dom.sz=670764&dom.img=71&dom.img.ext=29&dom.img.uniq=23&dom.script=12&dom.script.ext=5&dom.iframe=5&dom.iframe.ext=2&dom.link=108&dom.link.css=100&sb=1%0Abeacon%09'
```



## LICENSE 

MIT