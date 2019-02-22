package main

import (
	"fmt"
	"regexp"
)

func main()  {
	tmp:="abc a9c a7c ahc cvb cbv"
	reg1:=regexp.MustCompile("a.c")//result= [[abc] [a9c] [a7c] [ahc]]
	reg1=regexp.MustCompile("a[0-9]c")//result= [[a9c] [a7c]]
	if reg1==nil {
		fmt.Println("error")
		return
	}
	result:=reg1.FindAllStringSubmatch(tmp,-1)//-1表示全部解析
	fmt.Println("result=",result)
	tmp2:="1.22 3.45 3.23 7. .3 ahf i3ji"
	reg:=regexp.MustCompile(`\d+\.\d+`)//+表示匹配前面的字符一个或多个
	if reg==nil{
		fmt.Println("error")
		return
	}

	s2:=reg.FindAllStringSubmatch(tmp2,-1)
	fmt.Println("s2=",s2)
	tmp3:=`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
<script>window.NREUM||(NREUM={});NREUM.info={"beacon":"bam.nr-data.net","errorBeacon":"bam.nr-data.net","licenseKey":"2a000a1c4d","applicationID":"3221391","transactionName":"ewpeFRRWXVQDFxcSBUoMQBUVFlhWAgBA","queueTime":0,"applicationTime":78,"agent":""}</script>
<script>window.NREUM||(NREUM={}),__nr_require=function(e,n,t){function r(t){if(!n[t]){var o=n[t]={exports:{}};e[t][0].call(o.exports,function(n){var o=e[t][1][n];return r(o||n)},o,o.exports)}return n[t].exports}if("function"==typeof __nr_require)return __nr_require;for(var o=0;o<t.length;o++)r(t[o]);return r}({1:[function(e,n,t){function r(){}function o(e,n,t){return function(){return i(e,[c.now()].concat(u(arguments)),n?null:this,t),n?void 0:this}}var i=e("handle"),a=e(3),u=e(4),f=e("ee").get("tracer"),c=e("loader"),s=NREUM;"undefined"==typeof window.newrelic&&(newrelic=s);var p=["setPageViewName","setCustomAttribute","setErrorHandler","finished","addToTrace","inlineHit","addRelease"],d="api-",l=d+"ixn-";a(p,function(e,n){s[n]=o(d+n,!0,"api")}),s.addPageAction=o(d+"addPageAction",!0),s.setCurrentRouteName=o(d+"routeName",!0),n.exports=newrelic,s.interaction=function(){return(new r).get()};var m=r.prototype={createTracer:function(e,n){var t={},r=this,o="function"==typeof n;return i(l+"tracer",[c.now(),e,t],r),function(){if(f.emit((o?"":"no-")+"fn-start",[c.now(),r,o],t),o)try{return n.apply(this,arguments)}catch(e){throw f.emit("fn-err",[arguments,this,e],t),e}finally{f.emit("fn-end",[c.now()],t)}}}};a("actionText,setName,setAttribute,save,ignore,onEnd,getContext,end,get".split(","),function(e,n){m[n]=o(l+n)}),newrelic.noticeError=function(e,n){"string"==typeof e&&(e=new Error(e)),i("err",[e,c.now(),!1,n])}},{}],2:[function(e,n,t){function r(e,n){if(!o)return!1;if(e!==o)return!1;if(!n)return!0;if(!i)return!1;for(var t=i.split("."),r=n.split("."),a=0;a<r.length;a++)if(r[a]!==t[a])return!1;return!0}var o=null,i=null,a=/Version\/(\S+)\s+Safari/;if(navigator.userAgent){var u=navigator.userAgent,f=u.match(a);f&&u.indexOf("Chrome")===-1&&u.indexOf("Chromium")===-1&&(o="Safari",i=f[1])}n.exports={agent:o,version:i,match:r}},{}],3:[function(e,n,t){function r(e,n){var t=[],r="",i=0;for(r in e)o.call(e,r)&&(t[i]=n(r,e[r]),i+=1);return t}var o=Object.prototype.hasOwnProperty;n.exports=r},{}],4:[function(e,n,t){function r(e,n,t){n||(n=0),"undefined"==typeof t&&(t=e?e.length:0);for(var r=-1,o=t-n||0,i=Array(o<0?0:o);++r<o;)i[r]=e[n+r];return i}n.exports=r},{}],5:[function(e,n,t){n.exports={exists:"undefined"!=typeof window.performance&&window.performance.timing&&"undefined"!=typeof window.performance.timing.navigationStart}},{}],ee:[function(e,n,t){function r(){}function o(e){function n(e){return e&&e instanceof r?e:e?f(e,u,i):i()}function t(t,r,o,i){if(!d.aborted||i){e&&e(t,r,o);for(var a=n(o),u=v(t),f=u.length,c=0;c<f;c++)u[c].apply(a,r);var p=s[y[t]];return p&&p.push([b,t,r,a]),a}}function l(e,n){h[e]=v(e).concat(n)}function m(e,n){var t=h[e];if(t)for(var r=0;r<t.length;r++)t[r]===n&&t.splice(r,1)}function v(e){return h[e]||[]}function g(e){return p[e]=p[e]||o(t)}function w(e,n){c(e,function(e,t){n=n||"feature",y[t]=n,n in s||(s[n]=[])})}var h={},y={},b={on:l,addEventListener:l,removeEventListener:m,emit:t,get:g,listeners:v,context:n,buffer:w,abort:a,aborted:!1};return b}function i(){return new r}function a(){(s.api||s.feature)&&(d.aborted=!0,s=d.backlog={})}var u="nr@context",f=e("gos"),c=e(3),s={},p={},d=n.exports=o();d.backlog=s},{}],gos:[function(e,n,t){function r(e,n,t){if(o.call(e,n))return e[n];var r=t();if(Object.defineProperty&&Object.keys)try{return Object.defineProperty(e,n,{value:r,writable:!0,enumerable:!1}),r}catch(i){}return e[n]=r,r}var o=Object.prototype.hasOwnProperty;n.exports=r},{}],handle:[function(e,n,t){function r(e,n,t,r){o.buffer([e],r),o.emit(e,n,t)}var o=e("ee").get("handle");n.exports=r,r.ee=o},{}],id:[function(e,n,t){function r(e){var n=typeof e;return!e||"object"!==n&&"function"!==n?-1:e===window?0:a(e,i,function(){return o++})}var o=1,i="nr@id",a=e("gos");n.exports=r},{}],loader:[function(e,n,t){function r(){if(!E++){var e=x.info=NREUM.info,n=l.getElementsByTagName("script")[0];if(setTimeout(s.abort,3e4),!(e&&e.licenseKey&&e.applicationID&&n))return s.abort();c(y,function(n,t){e[n]||(e[n]=t)}),f("mark",["onload",a()+x.offset],null,"api");var t=l.createElement("script");t.src="https://"+e.agent,n.parentNode.insertBefore(t,n)}}function o(){"complete"===l.readyState&&i()}function i(){f("mark",["domContent",a()+x.offset],null,"api")}function a(){return O.exists&&performance.now?Math.round(performance.now()):(u=Math.max((new Date).getTime(),u))-x.offset}var u=(new Date).getTime(),f=e("handle"),c=e(3),s=e("ee"),p=e(2),d=window,l=d.document,m="addEventListener",v="attachEvent",g=d.XMLHttpRequest,w=g&&g.prototype;NREUM.o={ST:setTimeout,SI:d.setImmediate,CT:clearTimeout,XHR:g,REQ:d.Request,EV:d.Event,PR:d.Promise,MO:d.MutationObserver};var h=""+location,y={beacon:"bam.nr-data.net",errorBeacon:"bam.nr-data.net",agent:"js-agent.newrelic.com/nr-1118.min.js"},b=g&&w&&w[m]&&!/CriOS/.test(navigator.userAgent),x=n.exports={offset:u,now:a,origin:h,features:{},xhrWrappable:b,userAgent:p};e(1),l[m]?(l[m]("DOMContentLoaded",i,!1),d[m]("load",r,!1)):(l[v]("onreadystatechange",o),d[v]("onload",r)),f("mark",["firstbyte",u],null,"api");var E=0,O=e(5)},{}]},{},["loader"]);</script>
  <title>用户脚本</title>
  <meta name="description" value="针对您所访问的网站添加功能或解决问题的用户脚本。">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" media="all" href="/assets/application-169e09158f03e6c9bef1ee40fdcb9752f9ba2c5357542c9611c156ca4f388478.css" />
  <script src="/assets/application-95c24ada7b5f8059b0c0984c32d6ed7f3a6595fd9e595801bccae13c5c15821e.js"></script>
  <meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="G0+nDksPJXGLMrYZ1fHFE6nRR8RbElj+F6JIFb4S5+Hicswvo9R4jBVeVvtJphIffTEXNlpxaCfA4c1G/CwAGQ==" />
  <link rel="canonical" href="https://greasyfork.org/zh-CN/scripts">
  <link rel="icon" href="/assets/blacklogo16-bc64b9f7afdc9be4cbfa58bdd5fc2e5c098ad4bca3ad513a27b15602083fd5bc.png">
  <link href='https://fonts.googleapis.com/css?family=Open+Sans' rel='stylesheet' type='text/css'>
    <link rel="alternate" hreflang="x-default" href="/scripts">
      <link rel="alternate" hreflang="ar" href="/ar/scripts">
      <link rel="alternate" hreflang="bg" href="/bg/scripts">
      <link rel="alternate" hreflang="cs" href="/cs/scripts">
      <link rel="alternate" hreflang="da" href="/da/scripts">
      <link rel="alternate" hreflang="de" href="/de/scripts">
      <link rel="alternate" hreflang="el" href="/el/scripts">
      <link rel="alternate" hreflang="en" href="/en/scripts">
      <link rel="alternate" hreflang="es" href="/es/scripts">
      <link rel="alternate" hreflang="fi" href="/fi/scripts">
      <link rel="alternate" hreflang="fr" href="/fr/scripts">
      <link rel="alternate" hreflang="fr-CA" href="/fr-CA/scripts">
      <link rel="alternate" hreflang="he" href="/he/scripts">
      <link rel="alternate" hreflang="hu" href="/hu/scripts">
      <link rel="alternate" hreflang="id" href="/id/scripts">
      <link rel="alternate" hreflang="it" href="/it/scripts">
      <link rel="alternate" hreflang="ja" href="/ja/scripts">
      <link rel="alternate" hreflang="ko" href="/ko/scripts">
      <link rel="alternate" hreflang="nb" href="/nb/scripts">
      <link rel="alternate" hreflang="nl" href="/nl/scripts">
      <link rel="alternate" hreflang="pl" href="/pl/scripts">
      <link rel="alternate" hreflang="pt-BR" href="/pt-BR/scripts">
      <link rel="alternate" hreflang="ro" href="/ro/scripts">
      <link rel="alternate" hreflang="ru" href="/ru/scripts">
      <link rel="alternate" hreflang="sk" href="/sk/scripts">
      <link rel="alternate" hreflang="sv" href="/sv/scripts">
      <link rel="alternate" hreflang="th" href="/th/scripts">
      <link rel="alternate" hreflang="tr" href="/tr/scripts">
      <link rel="alternate" hreflang="uk" href="/uk/scripts">
      <link rel="alternate" hreflang="vi" href="/vi/scripts">
      <link rel="alternate" hreflang="zh-CN" href="/zh-CN/scripts">
      <link rel="alternate" hreflang="zh-TW" href="/zh-TW/scripts">

      <link rel="alternate" type="application/atom+xml" href="/zh-CN/scripts.atom?sort=created" title="最近创建的脚本">
      <link rel="alternate" type="application/atom+xml" href="/zh-CN/scripts.atom?sort=updated" title="最近更新的脚本">
      <link rel="alternate" type="application/json" href="/zh-CN/scripts.json" >
      <link rel="alternate" type="application/javascript" href="/zh-CN/scripts.jsonp?callback=callback" >
      <link rel="alternate" type="application/json" href="/zh-CN/scripts.json?meta=1" >
      <link rel="alternate" type="application/javascript" href="/zh-CN/scripts.jsonp?callback=callback&amp;meta=1" >

  <link rel="search" href="/zh-CN/opensearch.xml" type="application/opensearchdescription+xml" title="Greasy Fork 搜索" hreflang="zh-CN">

  
</head>
<body>

  <header id="main-header">
    <div class="width-constraint">
      <div id="site-name">
        <a href="/zh-CN"><img src="/assets/blacklogo96-e0c2c76180916332b7516ad47e1e206b42d131d36ff4afe98da3b1ba61fd5d6c.png" alt=""></a>
        <div id="site-name-text">
          <h1><a href="/zh-CN">Greasy Fork</a></h1>
        </div>
      </div>
      <div id="site-nav">
        <div id="nav-user-info">
            <span class="sign-in-link"><a rel="nofollow" href="/zh-CN/users/sign_in?return_to=%2Fzh-CN%2Fscripts">登录</a></span>

          <form id="language-selector" action="/scripts">
            <select id="language-selector-locale" name="locale">
                  <option data-language-url="/ar/scripts" value="ar">
                    العَرَبِيةُ (ar)
                  </option>
                  <option data-language-url="/bg/scripts" value="bg">
                    Български (bg)
                  </option>
                  <option data-language-url="/cs/scripts" value="cs">
                    Čeština (cs)
                  </option>
                  <option data-language-url="/da/scripts" value="da">
                    Dansk (da)
                  </option>
                  <option data-language-url="/de/scripts" value="de">
                    Deutsch (de)
                  </option>
                  <option data-language-url="/el/scripts" value="el">
                    Ελληνικά (el)
                  </option>
                  <option data-language-url="/en/scripts" value="en">
                    English (en)
                  </option>
                  <option data-language-url="/es/scripts" value="es">
                    Español (es)
                  </option>
                  <option data-language-url="/fi/scripts" value="fi">
                    Suomi (fi)
                  </option>
                  <option data-language-url="/fr/scripts" value="fr">
                    Français (fr)
                  </option>
                  <option data-language-url="/fr-CA/scripts" value="fr-CA">
                    Français canadien (fr-CA)
                  </option>
                  <option data-language-url="/he/scripts" value="he">
                    עברית (he)
                  </option>
                  <option data-language-url="/hu/scripts" value="hu">
                    Magyar (hu)
                  </option>
                  <option data-language-url="/id/scripts" value="id">
                    Bahasa Indonesia (id)
                  </option>
                  <option data-language-url="/it/scripts" value="it">
                    Italiano (it)
                  </option>
                  <option data-language-url="/ja/scripts" value="ja">
                    日本語 (ja)
                  </option>
                  <option data-language-url="/ko/scripts" value="ko">
                    한국어 (ko)
                  </option>
                  <option data-language-url="/nb/scripts" value="nb">
                    Bokmål (nb)
                  </option>
                  <option data-language-url="/nl/scripts" value="nl">
                    Nederlands (nl)
                  </option>
                  <option data-language-url="/pl/scripts" value="pl">
                    Polski (pl)
                  </option>
                  <option data-language-url="/pt-BR/scripts" value="pt-BR">
                    Português do Brasil (pt-BR)
                  </option>
                  <option data-language-url="/ro/scripts" value="ro">
                    Română (ro)
                  </option>
                  <option data-language-url="/ru/scripts" value="ru">
                    Русский (ru)
                  </option>
                  <option data-language-url="/sk/scripts" value="sk">
                    Slovenčina (sk)
                  </option>
                  <option data-language-url="/sv/scripts" value="sv">
                    Svenska (sv)
                  </option>
                  <option data-language-url="/th/scripts" value="th">
                    ภาษาไทย (th)
                  </option>
                  <option data-language-url="/tr/scripts" value="tr">
                    Türkçe (tr)
                  </option>
                  <option data-language-url="/uk/scripts" value="uk">
                    Українська (uk)
                  </option>
                  <option data-language-url="/vi/scripts" value="vi">
                    Tiếng Việt (vi)
                  </option>
                  <option data-language-url="/zh-CN/scripts" value="zh-CN" selected>
                    简体中文 (zh-CN)
                  </option>
                  <option data-language-url="/zh-TW/scripts" value="zh-TW">
                    繁體中文 (zh-TW)
                  </option>
              <option value="help">Help us translate!</option>
            </select><input id="language-selector-submit" type="submit" value="→">
            <script>
              /* submit is handled by js if enabled */
              document.getElementById("language-selector-submit").style.display = "none"
            </script>
          </form>
        </div>
        <nav>
          <li class="scripts-index-link"><a href="/zh-CN/scripts">脚本列表</a></li>
            <li class="forum-link"><a href="/zh-CN/forum/">论坛</a></li>
          <li class="help-link"><a href="/zh-CN/help">站点帮助</a></li>
          <li class="with-submenu">
            <a href="#" onclick="return false">更多</a>
            <nav>
              <li><a href="/zh-CN/search">高级搜索</a></li>
              <li><a href="/zh-CN/users">用户列表</a></li>
              <li><a href="/zh-CN/scripts/libraries">库</a></li>
              <li><a href="/zh-CN/moderator_actions">管理日志</a></li>
            </nav>
          </li>
        </nav>
      </div>
    </div>
  </header>

  <div class="width-constraint">
      <p class="notice"><b><a href="/en/scripts">Greasy Fork is available in English.</a></b></p>

    	
<div class="sidebarred">
  <div class="sidebarred-main-content">
    <div class="open-sidebar sidebar-collapsed">
      ☰
    </div>



      <p>现在显示所有语言的结果。<a href="/zh-CN/scripts?filter_locale=1">只显示 简体中文 的结果</a></p>

      <ol id="browse-script-list" class="script-list">
        <li data-script-id="39504" data-script-name="百度网盘直接下载助手 直链加速版" data-script-author-id="174924" data-script-author-name="syhyz1990" data-script-daily-installs="4665" data-script-total-installs="1001305" data-script-rating-score="98.1" data-script-created-date="2018-03-15" data-script-updated-date="2019-02-20" data-script-type="public" data-script-version="2.0.1" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/39504-%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD%E5%8A%A9%E6%89%8B-%E7%9B%B4%E9%93%BE%E5%8A%A0%E9%80%9F%E7%89%88">百度网盘直接下载助手 直链加速版</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				百度网盘高速下载 支持IDM  [2019-02-20] 修复我的网盘API下载返回403的问题
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/174924-syhyz1990">syhyz1990</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>4,665</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>1,001,305</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.1"><span><span class="good-rating-count" title="好评或收藏的人数。">1440</span>
<span class="ok-rating-count" title="评级为一般的人数。">19</span>
<span class="bad-rating-count" title="评级为差评的人数。">8</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-03-15T05:34:34+00:00">2018-03-15</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-20T11:44:01+00:00">2019-02-20</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="370005" data-script-name="一键VIP视频解析、去广告（全网）,一站式音乐搜索下载 2019-02-18 更新，报错请及时反馈" data-script-author-id="189391" data-script-author-name="Max zhang" data-script-daily-installs="2154" data-script-total-installs="278390" data-script-rating-score="96.7" data-script-created-date="2018-07-03" data-script-updated-date="2019-02-18" data-script-type="public" data-script-version="3.1.7" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/370005-%E4%B8%80%E9%94%AEvip%E8%A7%86%E9%A2%91%E8%A7%A3%E6%9E%90-%E5%8E%BB%E5%B9%BF%E5%91%8A-%E5%85%A8%E7%BD%91-%E4%B8%80%E7%AB%99%E5%BC%8F%E9%9F%B3%E4%B9%90%E6%90%9C%E7%B4%A2%E4%B8%8B%E8%BD%BD-2019-02-18-%E6%9B%B4%E6%96%B0-%E6%8A%A5%E9%94%99%E8%AF%B7%E5%8F%8A%E6%97%B6%E5%8F%8D%E9%A6%88">一键VIP视频解析、去广告（全网）,一站式音乐搜索下载 2019-02-18 更新，报错请及时反馈</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				在视频播放页悬浮VIP按钮，可在线播放vip视频；支持优酷vip，腾讯vip，爱奇艺vip，芒果vip，乐视vip等常用视频...一站式音乐搜索解决方案，网易云音乐，QQ音乐，酷狗音乐，酷我音乐，虾米音乐，百度音乐，蜻蜓FM，荔枝FM，喜马拉雅，优惠券查询
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/189391-max-zhang">Max zhang</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>2,154</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>278,390</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="96.7"><span><span class="good-rating-count" title="好评或收藏的人数。">320</span>
<span class="ok-rating-count" title="评级为一般的人数。">3</span>
<span class="bad-rating-count" title="评级为差评的人数。">3</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-07-03T01:48:34+00:00">2018-07-03</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-18T06:16:13+00:00">2019-02-18</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="370634" data-script-name="懒人专用，全网VIP视频破解去广告免费看，百度网盘直接下载、知乎视频下载、咸鱼搜索框等多合一版。2019全年更新，放心使用。" data-script-author-id="198522" data-script-author-name="懒蛤蛤" data-script-daily-installs="1800" data-script-total-installs="104052" data-script-rating-score="97.4" data-script-created-date="2018-07-27" data-script-updated-date="2019-02-20" data-script-type="public" data-script-version="1.2.15" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/370634-%E6%87%92%E4%BA%BA%E4%B8%93%E7%94%A8-%E5%85%A8%E7%BD%91vip%E8%A7%86%E9%A2%91%E7%A0%B4%E8%A7%A3%E5%8E%BB%E5%B9%BF%E5%91%8A%E5%85%8D%E8%B4%B9%E7%9C%8B-%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD-%E7%9F%A5%E4%B9%8E%E8%A7%86%E9%A2%91%E4%B8%8B%E8%BD%BD-%E5%92%B8%E9%B1%BC%E6%90%9C%E7%B4%A2%E6%A1%86%E7%AD%89%E5%A4%9A%E5%90%88%E4%B8%80%E7%89%88-2019%E5%85%A8%E5%B9%B4%E6%9B%B4%E6%96%B0-%E6%94%BE%E5%BF%83%E4%BD%BF%E7%94%A8">懒人专用，全网VIP视频破解去广告免费看，百度网盘直接下载、知乎视频下载、咸鱼搜索框等多合一版。2019全年更新，放心使用。</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				自用组合型多功能脚本，优酷、爱奇艺、腾讯等全网VIP视频破解去广告免费看，百度网盘直接下载，知乎视频下载，闲鱼搜索框，优惠券领取等几个自己常用的功能。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/198522-%E6%87%92%E8%9B%A4%E8%9B%A4">懒蛤蛤</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>1,800</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>104,052</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="97.4"><span><span class="good-rating-count" title="好评或收藏的人数。">216</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-07-27T02:35:32+00:00">2018-07-27</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-20T11:33:47+00:00">2019-02-20</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="38857" data-script-name="ScriptSource: The Leading Portal for Online Extensions &amp; Apps [Diservers / MooMoo.io / Krunker.io]" data-script-author-id="171504" data-script-author-name="Sam-DevZ" data-script-daily-installs="1513" data-script-total-installs="154393" data-script-rating-score="68.6" data-script-created-date="2018-02-24" data-script-updated-date="2019-02-21" data-script-type="public" data-script-version="25.2" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/38857-scriptsource-the-leading-portal-for-online-extensions-apps-diservers-moomoo-io-krunker-io">ScriptSource: The Leading Portal for Online Extensions &amp; Apps [Diservers / MooMoo.io / Krunker.io]</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Currently trusted by over 100,000 users!
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/171504-sam-devz">Sam-DevZ</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>1,513</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>154,393</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="68.6"><span><span class="good-rating-count" title="好评或收藏的人数。">104</span>
<span class="ok-rating-count" title="评级为一般的人数。">20</span>
<span class="bad-rating-count" title="评级为差评的人数。">26</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-02-24T19:53:18+00:00">2018-02-24</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-21T02:43:56+00:00">2019-02-21</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="14178" data-script-name="AC-baidu:重定向优化百度搜狗谷歌搜索_去广告_favicon_双列" data-script-author-id="18978" data-script-author-name="inDarkness" data-script-daily-installs="1486" data-script-total-installs="516074" data-script-rating-score="99.4" data-script-created-date="2015-11-24" data-script-updated-date="2019-02-16" data-script-type="public" data-script-version="23.12" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/14178-ac-baidu-%E9%87%8D%E5%AE%9A%E5%90%91%E4%BC%98%E5%8C%96%E7%99%BE%E5%BA%A6%E6%90%9C%E7%8B%97%E8%B0%B7%E6%AD%8C%E6%90%9C%E7%B4%A2-%E5%8E%BB%E5%B9%BF%E5%91%8A-favicon-%E5%8F%8C%E5%88%97">AC-baidu:重定向优化百度搜狗谷歌搜索_去广告_favicon_双列</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				1.繞過百度、搜狗、谷歌、好搜搜索結果中的自己的跳轉鏈接，直接訪問原始網頁-反正都能看懂 2.新增拦截百度百家号的无用推广数据-支持其他站点 3.去除百度的多余广告 4添加Favicon显示 5.页面CSS 6.添加计数 7.开关选择以上功能
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/18978-indarkness">inDarkness</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>1,486</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>516,074</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.4"><span><span class="good-rating-count" title="好评或收藏的人数。">2825</span>
<span class="ok-rating-count" title="评级为一般的人数。">9</span>
<span class="bad-rating-count" title="评级为差评的人数。">4</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2015-11-24T23:08:15+00:00">2015-11-24</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-16T13:34:15+00:00">2019-02-16</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="374052" data-script-name="KRUNKER.IO MODS,HACKS,CHEATS UNBLOCKED KRUNKERIO BY IO-MODS! GODMODE-AUTO AIM-ESP-AIMBOT" data-script-author-id="216564" data-script-author-name="IO-MODS" data-script-daily-installs="1179" data-script-total-installs="101540" data-script-rating-score="48.8" data-script-created-date="2018-11-06" data-script-updated-date="2018-11-24" data-script-type="public" data-script-version="1.5.3" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/374052-krunker-io-mods-hacks-cheats-unblocked-krunkerio-by-io-mods-godmode-auto-aim-esp-aimbot">KRUNKER.IO MODS,HACKS,CHEATS UNBLOCKED KRUNKERIO BY IO-MODS! GODMODE-AUTO AIM-ESP-AIMBOT</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				krunker.io hack,Auto Aim,Auto Heal,Auto Block,Auto Fire,Aimbot
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/216564-io-mods">IO-MODS</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>1,179</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>101,540</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="48.8"><span><span class="good-rating-count" title="好评或收藏的人数。">16</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">7</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-11-06T23:12:57+00:00">2018-11-06</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-11-24T19:35:59+00:00">2018-11-24</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="27530" data-script-name="破解VIP会员视频集合" data-script-author-id="104201" data-script-author-name="黄盐" data-script-daily-installs="1008" data-script-total-installs="1413290" data-script-rating-score="98.6" data-script-created-date="2017-02-20" data-script-updated-date="2018-08-14" data-script-type="public" data-script-version="4.2.5" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/27530-%E7%A0%B4%E8%A7%A3vip%E4%BC%9A%E5%91%98%E8%A7%86%E9%A2%91%E9%9B%86%E5%90%88">破解VIP会员视频集合</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				一键破解[优酷|腾讯|乐视|爱奇艺|芒果|AB站|音悦台]等VIP或会员视频，解析接口贵精不贵多，绝对够用。详细方法看说明和图片。包含了[一键VIP视频解析、去广告（全网） xxxx-xx-xx 可用▶mark zhang][VIP视频在线解析破解去广告(全网)xx.xx.xx更新可用▶sonimei134][破解全网VIP视频会员-去广告▶ttmsjx][VIP会员视频解析▶龙轩][酷绘-破解VIP会员视频▶ahuiabc2003]以及[VIP视频破解▶hoothin]的部分接口。[Tampermonkey | Violentmonkey | Greasymonkey 4.0+]
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/104201-%E9%BB%84%E7%9B%90">黄盐</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>1,008</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>1,413,290</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.6"><span><span class="good-rating-count" title="好评或收藏的人数。">2634</span>
<span class="ok-rating-count" title="评级为一般的人数。">27</span>
<span class="bad-rating-count" title="评级为差评的人数。">13</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-02-20T05:36:37+00:00">2017-02-20</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-08-14T15:34:30+00:00">2018-08-14</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="374877" data-script-name="Krunker.io 2019 Aimbot (Hacks,Mods,Cheats) | NO DISCONNECT ISSUE | KRUNKERIO.ORG" data-script-author-id="228916" data-script-author-name="Mr.Coder" data-script-daily-installs="905" data-script-total-installs="22147" data-script-rating-score="27.4" data-script-created-date="2018-11-28" data-script-updated-date="2019-02-04" data-script-type="public" data-script-version="5.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/374877-krunker-io-2019-aimbot-hacks-mods-cheats-no-disconnect-issue-krunkerio-org">Krunker.io 2019 Aimbot (Hacks,Mods,Cheats) | NO DISCONNECT ISSUE | KRUNKERIO.ORG</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Aimbot, Unlimited Ammo, Auto Heal, ESP, Wall Hack, Unlimited Ammo... -2019 krunkerio hack- ADBLOCK
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/228916-mr-coder">Mr.Coder</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>905</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>22,147</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="27.4"><span><span class="good-rating-count" title="好评或收藏的人数。">10</span>
<span class="ok-rating-count" title="评级为一般的人数。">6</span>
<span class="bad-rating-count" title="评级为差评的人数。">14</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-11-28T03:52:07+00:00">2018-11-28</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-04T07:47:31+00:00">2019-02-04</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="23635" data-script-name="百度网盘直接下载助手" data-script-author-id="70373" data-script-author-name="ivesjay" data-script-daily-installs="835" data-script-total-installs="1782516" data-script-rating-score="98.3" data-script-created-date="2016-10-01" data-script-updated-date="2017-05-26" data-script-type="public" data-script-version="0.9.24" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/23635-%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD%E5%8A%A9%E6%89%8B">百度网盘直接下载助手</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				直接下载百度网盘和百度网盘分享的文件,避免下载文件时调用百度网盘客户端,获取网盘文件的直接下载地址
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/70373-ivesjay">ivesjay</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>835</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>1,782,516</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.3"><span><span class="good-rating-count" title="好评或收藏的人数。">3482</span>
<span class="ok-rating-count" title="评级为一般的人数。">37</span>
<span class="bad-rating-count" title="评级为差评的人数。">27</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-10-01T10:07:17+00:00">2016-10-01</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2017-05-26T06:06:51+00:00">2017-05-26</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="26638" data-script-name="EX-百度云盘" data-script-author-id="87569" data-script-author-name="gxvv" data-script-daily-installs="745" data-script-total-installs="620176" data-script-rating-score="98.6" data-script-created-date="2017-01-18" data-script-updated-date="2018-08-20" data-script-type="public" data-script-version="0.3.3" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/26638-ex-%E7%99%BE%E5%BA%A6%E4%BA%91%E7%9B%98">EX-百度云盘</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				[下载大文件] [批量下载] [文件夹下载] [百度网盘] [百度云盘] [企业版]
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/87569-gxvv">gxvv</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>745</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>620,176</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.6"><span><span class="good-rating-count" title="好评或收藏的人数。">1339</span>
<span class="ok-rating-count" title="评级为一般的人数。">8</span>
<span class="bad-rating-count" title="评级为差评的人数。">6</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-01-18T07:02:39+00:00">2017-01-18</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-08-20T06:20:04+00:00">2018-08-20</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="28497" data-script-name="网页限制解除(改)" data-script-author-id="106222" data-script-author-name="qxin i" data-script-daily-installs="749" data-script-total-installs="246123" data-script-rating-score="98.8" data-script-created-date="2017-03-28" data-script-updated-date="2018-07-02" data-script-type="public" data-script-version="4.1.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/28497-remove-web-limits-modified">网页限制解除(改)</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				通杀大部分网站,可以解除禁止复制、剪切、选择文本、右键菜单的限制。原作者cat73,因为和搜索跳转脚本冲突,遂进行了改动,改为黑名单制。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/106222-qxin-i">qxin i</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>749</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>246,123</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.8"><span><span class="good-rating-count" title="好评或收藏的人数。">1167</span>
<span class="ok-rating-count" title="评级为一般的人数。">5</span>
<span class="bad-rating-count" title="评级为差评的人数。">4</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-03-28T05:48:53+00:00">2017-03-28</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-07-02T14:32:31+00:00">2018-07-02</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="376712" data-script-name="百度文库文字复制及原格式下载" data-script-author-id="235134" data-script-author-name="刚仔文库" data-script-daily-installs="711" data-script-total-installs="20811" data-script-rating-score="90.0" data-script-created-date="2019-01-15" data-script-updated-date="2019-02-17" data-script-type="public" data-script-version="0.27" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/376712-%E7%99%BE%E5%BA%A6%E6%96%87%E5%BA%93%E6%96%87%E5%AD%97%E5%A4%8D%E5%88%B6%E5%8F%8A%E5%8E%9F%E6%A0%BC%E5%BC%8F%E4%B8%8B%E8%BD%BD">百度文库文字复制及原格式下载</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				为百度文库页面添加两个按钮，一键复制百度文库文字,一键加群获取原格式文档
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/235134-%E5%88%9A%E4%BB%94%E6%96%87%E5%BA%93">刚仔文库</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>711</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>20,811</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="90.0"><span><span class="good-rating-count" title="好评或收藏的人数。">73</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">2</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2019-01-15T01:30:27+00:00">2019-01-15</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-17T04:56:16+00:00">2019-02-17</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="329484" data-script-name="豆瓣资源下载大师：1秒搞定豆瓣电影|音乐|图书下载" data-script-author-id="4458" data-script-author-name="bimzcy" data-script-daily-installs="645" data-script-total-installs="128137" data-script-rating-score="98.7" data-script-created-date="2018-05-06" data-script-updated-date="2019-02-20" data-script-type="public" data-script-version="7.1.9" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/329484-%E8%B1%86%E7%93%A3%E8%B5%84%E6%BA%90%E4%B8%8B%E8%BD%BD%E5%A4%A7%E5%B8%88-1%E7%A7%92%E6%90%9E%E5%AE%9A%E8%B1%86%E7%93%A3%E7%94%B5%E5%BD%B1-%E9%9F%B3%E4%B9%90-%E5%9B%BE%E4%B9%A6%E4%B8%8B%E8%BD%BD">豆瓣资源下载大师：1秒搞定豆瓣电影|音乐|图书下载</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				【装这一个脚本就够了！可能是你遇到的最好的豆瓣增强脚本】聚合数百家资源网站，通过右侧边栏1秒告诉你哪些网站能下载豆瓣页面上的电影|电视剧|纪录片|综艺|动画|音乐|图书等，有资源的网站显示绿色，没资源的网站显示黄色，就这么直观！所有豆瓣条目均提供在线播放|阅读、字幕|歌词下载及PT|NZB|BT|磁力|百度盘|115网盘等下载链接，加入官网打死也不出的豆列搜索功能，此外还能给豆瓣条目额外添加IMDB评分|IMDB TOP 250|Metascore评分|烂番茄评分|AniDB评分|Bgm评分|MAL|亚马逊评分等更多评分形式。官方电报群：@doubandown，官方QQ群：691023412，验证口令：imdb
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/4458-bimzcy">bimzcy</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>645</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>128,137</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.7"><span><span class="good-rating-count" title="好评或收藏的人数。">771</span>
<span class="ok-rating-count" title="评级为一般的人数。">4</span>
<span class="bad-rating-count" title="评级为差评的人数。">2</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-05-06T00:47:42+00:00">2018-05-06</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-20T04:38:37+00:00">2019-02-20</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="29762" data-script-name="网盘自动填写密码【威力加强版】" data-script-author-id="3128" data-script-author-name="极品小猫" data-script-daily-installs="566" data-script-total-installs="201873" data-script-rating-score="99.6" data-script-created-date="2017-05-15" data-script-updated-date="2019-02-17" data-script-type="public" data-script-version="3.9.2.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/29762-%E7%BD%91%E7%9B%98%E8%87%AA%E5%8A%A8%E5%A1%AB%E5%86%99%E5%AF%86%E7%A0%81-%E5%A8%81%E5%8A%9B%E5%8A%A0%E5%BC%BA%E7%89%88">网盘自动填写密码【威力加强版】</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				智能融合网盘密码到网址中，打开网盘链接时不再需要手动复制密码，并自动提交密码，一路畅通无阻。同时记录网盘信息，当你再次打开该分享文件时，不再需要去找提取码，同时可追溯网盘地址的来源。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/3128-%E6%9E%81%E5%93%81%E5%B0%8F%E7%8C%AB">极品小猫</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>566</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>201,873</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.6"><span><span class="good-rating-count" title="好评或收藏的人数。">1257</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-05-15T03:50:01+00:00">2017-05-15</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-17T09:03:45+00:00">2019-02-17</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="25718" data-script-name="解除B站区域限制" data-script-author-id="86730" data-script-author-name="ipcjs" data-script-daily-installs="557" data-script-total-installs="416266" data-script-rating-score="98.8" data-script-created-date="2016-12-16" data-script-updated-date="2019-02-16" data-script-type="public" data-script-version="7.3.2" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/25718-%E8%A7%A3%E9%99%A4b%E7%AB%99%E5%8C%BA%E5%9F%9F%E9%99%90%E5%88%B6">解除B站区域限制</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				通过替换获取视频地址接口的方式, 实现解除B站区域限制; 只对HTML5播放器生效;
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/86730-ipcjs">ipcjs</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>557</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>416,266</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.8"><span><span class="good-rating-count" title="好评或收藏的人数。">1160</span>
<span class="ok-rating-count" title="评级为一般的人数。">7</span>
<span class="bad-rating-count" title="评级为差评的人数。">3</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-12-16T04:22:47+00:00">2016-12-16</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-16T14:15:19+00:00">2019-02-16</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="369418" data-script-name="The best Youtube Downloader, download video MP4, AVI, MP3, HD, 1080P, 2K, 4k &amp; 8K" data-script-author-id="190574" data-script-author-name="hello123" data-script-daily-installs="521" data-script-total-installs="104272" data-script-rating-score="87.5" data-script-created-date="2018-06-11" data-script-updated-date="2018-11-07" data-script-type="public" data-script-version="5.1" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/369418-the-best-youtube-downloader-download-video-mp4-avi-mp3-hd-1080p-2k-4k-8k">The best Youtube Downloader, download video MP4, AVI, MP3, HD, 1080P, 2K, 4k &amp; 8K</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Download any video and music (audio) from Youtube, Twitter, Vimeo, Facebook, Instagram, SoundCloud, Dailymotion, Liveleak, Break, Imgur, Mashable, Reddit, 1TV, 9gag, VK, TED, youku, bilibili, IMDb, ESPN, Flickr, Bandcamp, pornhub, 9gag, VK.com, ok.ru, tv.com and 10,000 more sites for free. Also support to download subtitles. Free, fast and easy to use. No need to install any annoying softwares. Supporting MP4, WEBM, AVI, 3GP, FLV, H64, ACC, FLA, MP3, M4A, 8K, 6K,4K, 2K, 1080, 720, 480, 360, etc.
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/190574-hello123">hello123</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>521</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>104,272</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="87.5"><span><span class="good-rating-count" title="好评或收藏的人数。">223</span>
<span class="ok-rating-count" title="评级为一般的人数。">5</span>
<span class="bad-rating-count" title="评级为差评的人数。">18</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-06-11T10:20:59+00:00">2018-06-11</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-11-07T06:42:40+00:00">2018-11-07</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="30545" data-script-name="视频站启用html5播放器" data-script-author-id="7036" data-script-author-name="xinggsf" data-script-daily-installs="485" data-script-total-installs="298280" data-script-rating-score="98.9" data-script-created-date="2017-06-12" data-script-updated-date="2018-12-24" data-script-type="public" data-script-version="1.2.8" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/30545-%E8%A7%86%E9%A2%91%E7%AB%99%E5%90%AF%E7%94%A8html5%E6%92%AD%E6%94%BE%E5%99%A8">视频站启用html5播放器</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				三大功能 。启用html5播放器；万能网页全屏；添加快捷键：快进、快退、暂停/播放、音量、下一集、切换(网页)全屏、上下帧、播放速度。支持视频站点：油管、TED、优.土、QQ、B站、芒果TV、新浪、微博、网易[娱乐、云课堂、新闻]、搜狐、乐视、风行、百度云视频等；直播：斗鱼、熊猫、YY、虎牙、龙珠、战旗。可增加自定义站点
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/7036-xinggsf">xinggsf</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>485</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>298,280</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.9"><span><span class="good-rating-count" title="好评或收藏的人数。">1464</span>
<span class="ok-rating-count" title="评级为一般的人数。">7</span>
<span class="bad-rating-count" title="评级为差评的人数。">5</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-06-12T14:41:55+00:00">2017-06-12</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-12-24T08:55:31+00:00">2018-12-24</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="369221" data-script-name="百度云百度网盘直接下载助手直链加速版支持迅雷+VIP视频在线解析破解去广告+【一键领取】淘宝购物领取大额优惠券，购物立省80%！-2019.1.25加强稳定版" data-script-author-id="187972" data-script-author-name="淘优惠" data-script-daily-installs="442" data-script-total-installs="281865" data-script-rating-score="92.7" data-script-created-date="2018-06-04" data-script-updated-date="2019-01-24" data-script-type="public" data-script-version="10.27" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/369221-%E7%99%BE%E5%BA%A6%E4%BA%91%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD%E5%8A%A9%E6%89%8B%E7%9B%B4%E9%93%BE%E5%8A%A0%E9%80%9F%E7%89%88%E6%94%AF%E6%8C%81%E8%BF%85%E9%9B%B7-vip%E8%A7%86%E9%A2%91%E5%9C%A8%E7%BA%BF%E8%A7%A3%E6%9E%90%E7%A0%B4%E8%A7%A3%E5%8E%BB%E5%B9%BF%E5%91%8A-%E4%B8%80%E9%94%AE%E9%A2%86%E5%8F%96-%E6%B7%98%E5%AE%9D%E8%B4%AD%E7%89%A9%E9%A2%86%E5%8F%96%E5%A4%A7%E9%A2%9D%E4%BC%98%E6%83%A0%E5%88%B8-%E8%B4%AD%E7%89%A9%E7%AB%8B%E7%9C%8180-2019-1-25%E5%8A%A0%E5%BC%BA%E7%A8%B3%E5%AE%9A%E7%89%88">百度云百度网盘直接下载助手直链加速版支持迅雷+VIP视频在线解析破解去广告+【一键领取】淘宝购物领取大额优惠券，购物立省80%！-2019.1.25加强稳定版</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				修复百度云链接不能下载问题！！！百度云百度网盘获取直接下载链接+压缩下载链接；大文件/单文件/多文件/文件夹，支持点击直接下载免强制调用百度网盘客户端；通过淘宝客网站，查询商家设置的隐藏优惠券信息！直接领取优惠券购买。新增淘宝自动查券功能，给你购物带来全新体验。在视频标题旁上显示“点击VIP解析播放”按钮，在线播放vip视频；支持优酷vip，腾讯vip，爱奇艺vip，芒果vip，乐视vip等常用视频！新增一个全新的购物网站，备案号：浙ICP备17026363号，给您安全放心的购物环境！！！
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/187972-%E6%B7%98%E4%BC%98%E6%83%A0">淘优惠</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>442</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>281,865</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="92.7"><span><span class="good-rating-count" title="好评或收藏的人数。">258</span>
<span class="ok-rating-count" title="评级为一般的人数。">5</span>
<span class="bad-rating-count" title="评级为差评的人数。">9</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-06-04T13:21:02+00:00">2018-06-04</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-24T13:34:56+00:00">2019-01-24</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="377517" data-script-name="Krunker.io Aimbot" data-script-author-id="244978" data-script-author-name="GET16HAX" data-script-daily-installs="448" data-script-total-installs="5588" data-script-rating-score="52.9" data-script-created-date="2019-02-07" data-script-updated-date="2019-02-11" data-script-type="public" data-script-version="25" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/377517-krunker-io-aimbot">Krunker.io Aimbot</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Currently trusted by over 100,000 users! If it&#39;s not working wait for a update!
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/244978-get16hax">GET16HAX</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>448</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>5,588</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="52.9"><span><span class="good-rating-count" title="好评或收藏的人数。">6</span>
<span class="ok-rating-count" title="评级为一般的人数。">2</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2019-02-07T11:39:51+00:00">2019-02-07</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-11T09:58:42+00:00">2019-02-11</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="37058" data-script-name="全网音乐一键免费下载 一键搜索 在线试听 【高清品质请选择QQ下载源】2019.1.10更新" data-script-author-id="165139" data-script-author-name="村长" data-script-daily-installs="432" data-script-total-installs="200869" data-script-rating-score="98.4" data-script-created-date="2018-01-05" data-script-updated-date="2019-01-09" data-script-type="public" data-script-version="1.6.34" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/37058-%E5%85%A8%E7%BD%91%E9%9F%B3%E4%B9%90%E4%B8%80%E9%94%AE%E5%85%8D%E8%B4%B9%E4%B8%8B%E8%BD%BD-%E4%B8%80%E9%94%AE%E6%90%9C%E7%B4%A2-%E5%9C%A8%E7%BA%BF%E8%AF%95%E5%90%AC-%E9%AB%98%E6%B8%85%E5%93%81%E8%B4%A8%E8%AF%B7%E9%80%89%E6%8B%A9qq%E4%B8%8B%E8%BD%BD%E6%BA%90-2019-1-10%E6%9B%B4%E6%96%B0">全网音乐一键免费下载 一键搜索 在线试听 【高清品质请选择QQ下载源】2019.1.10更新</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				全网音乐在线试听 一键免费下载 一键搜索 提供多站合一 音乐试听 音乐下载 音乐搜索解决方案，网易云音乐，QQ音乐，酷狗音乐，酷我音乐，虾米音乐，百度音乐，蜻蜓FM，荔枝FM，喜马拉雅
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/165139-%E6%9D%91%E9%95%BF">村长</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>432</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>200,869</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.4"><span><span class="good-rating-count" title="好评或收藏的人数。">742</span>
<span class="ok-rating-count" title="评级为一般的人数。">7</span>
<span class="bad-rating-count" title="评级为差评的人数。">2</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-01-05T15:26:30+00:00">2018-01-05</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-09T16:20:19+00:00">2019-01-09</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="14466" data-script-name="购物党自动比价工具-领取淘宝内部券" data-script-author-id="54" data-script-author-name="文科" data-script-daily-installs="410" data-script-total-installs="282654" data-script-rating-score="99.4" data-script-created-date="2015-12-03" data-script-updated-date="2018-09-05" data-script-type="public" data-script-version="3.0.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/14466-%E8%B4%AD%E7%89%A9%E5%85%9A%E8%87%AA%E5%8A%A8%E6%AF%94%E4%BB%B7%E5%B7%A5%E5%85%B7-%E9%A2%86%E5%8F%96%E6%B7%98%E5%AE%9D%E5%86%85%E9%83%A8%E5%88%B8">购物党自动比价工具-领取淘宝内部券</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				［含有购物党的返利］浏览商品页面时，自动比较同款商品在淘宝/京东/亚马逊/当当/苏宁/等百家商城的最低价，提供价格历史、口碑评分等查询。支持商品促销活动，商城优惠信息查询，商品可全网收藏，降价提醒。支持链家、我爱我家、中原地产等主流房产网站房源价格走势查询，为买房人士提供决策参考。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/54-%E6%96%87%E7%A7%91">文科</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>410</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>282,654</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.4"><span><span class="good-rating-count" title="好评或收藏的人数。">1844</span>
<span class="ok-rating-count" title="评级为一般的人数。">6</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2015-12-03T13:28:10+00:00">2015-12-03</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-09-05T16:01:44+00:00">2018-09-05</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="30117" data-script-name="吾爱破解网盘激活&amp;提取码自动补全" data-script-author-id="47767" data-script-author-name="猎隼丶止戈" data-script-daily-installs="403" data-script-total-installs="133939" data-script-rating-score="99.3" data-script-created-date="2017-05-30" data-script-updated-date="2019-01-30" data-script-type="public" data-script-version="0.1.0.7" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/30117-%E5%90%BE%E7%88%B1%E7%A0%B4%E8%A7%A3%E7%BD%91%E7%9B%98%E6%BF%80%E6%B4%BB-%E6%8F%90%E5%8F%96%E7%A0%81%E8%87%AA%E5%8A%A8%E8%A1%A5%E5%85%A8">吾爱破解网盘激活&amp;提取码自动补全</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				激活吾爱破解论坛中的百度网盘链接，并自动补全提取码然后跳转到分享地址
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/47767-%E7%8C%8E%E9%9A%BC%E4%B8%B6%E6%AD%A2%E6%88%88">猎隼丶止戈</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>403</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>133,939</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.3"><span><span class="good-rating-count" title="好评或收藏的人数。">556</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-05-30T07:31:15+00:00">2017-05-30</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-30T08:40:06+00:00">2019-01-30</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="373914" data-script-name="91解析" data-script-author-id="150069" data-script-author-name="shopkeeperV" data-script-daily-installs="403" data-script-total-installs="29906" data-script-rating-score="98.3" data-script-created-date="2018-11-03" data-script-updated-date="2019-02-05" data-script-type="public" data-script-version="3.2.0" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/373914-91%E8%A7%A3%E6%9E%90">91解析</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				次数是无限的，但精力是有限的。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/150069-shopkeeperv">shopkeeperV</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>403</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>29,906</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.3"><span><span class="good-rating-count" title="好评或收藏的人数。">269</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-11-03T12:35:25+00:00">2018-11-03</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-05T04:37:11+00:00">2019-02-05</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="2600" data-script-name="跳过网站等待、验证码及登录" data-script-author-id="44" data-script-author-name="Jixun.Moe" data-script-daily-installs="369" data-script-total-installs="453292" data-script-rating-score="99.8" data-script-created-date="2014-06-17" data-script-updated-date="2017-04-28" data-script-type="public" data-script-version="4.0.720" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/2600-%E8%B7%B3%E8%BF%87%E7%BD%91%E7%AB%99%E7%AD%89%E5%BE%85-%E9%AA%8C%E8%AF%81%E7%A0%81%E5%8F%8A%E7%99%BB%E5%BD%95">跳过网站等待、验证码及登录</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				移除各类网站验证码、登录、倒计时及更多!
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/44-jixun-moe">Jixun.Moe</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>369</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>453,292</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.8"><span><span class="good-rating-count" title="好评或收藏的人数。">2954</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">2</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2014-06-17T22:33:57+00:00">2014-06-17</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2017-04-28T12:40:53+00:00">2017-04-28</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="24192" data-script-name="百度广告(首尾推广及右侧广告)清理" data-script-author-id="8227" data-script-author-name="hoothin" data-script-daily-installs="388" data-script-total-installs="219157" data-script-rating-score="99.3" data-script-created-date="2016-10-21" data-script-updated-date="2019-01-31" data-script-type="public" data-script-version="1.25" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/24192-%E7%99%BE%E5%BA%A6%E5%B9%BF%E5%91%8A-%E9%A6%96%E5%B0%BE%E6%8E%A8%E5%B9%BF%E5%8F%8A%E5%8F%B3%E4%BE%A7%E5%B9%BF%E5%91%8A-%E6%B8%85%E7%90%86">百度广告(首尾推广及右侧广告)清理</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				彻底清理百度搜索(www.baidu.com)结果首尾的推广广告、二次顽固广告、右侧广告与百家号信息，并防止反复
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/8227-hoothin">hoothin</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>388</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>219,157</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.3"><span><span class="good-rating-count" title="好评或收藏的人数。">846</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-10-21T13:21:19+00:00">2016-10-21</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-31T07:23:41+00:00">2019-01-31</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="369400" data-script-name="本地 YouTube 下载器" data-script-author-id="72365" data-script-author-name="3142 maple" data-script-daily-installs="332" data-script-total-installs="53330" data-script-rating-score="94.5" data-script-created-date="2018-06-10" data-script-updated-date="2019-01-31" data-script-type="public" data-script-version="0.6.10" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/369400-local-youtube-downloader">本地 YouTube 下载器</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				不需要透过第三方的服务就能下载 YouTube 影片。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/72365-3142-maple">3142 maple</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>332</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>53,330</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="94.5"><span><span class="good-rating-count" title="好评或收藏的人数。">225</span>
<span class="ok-rating-count" title="评级为一般的人数。">2</span>
<span class="bad-rating-count" title="评级为差评的人数。">5</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-06-10T14:48:08+00:00">2018-06-10</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-31T09:59:26+00:00">2019-01-31</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="372452" data-script-name="CSDN自动展开+去广告+净化剪贴板+免登陆" data-script-author-id="214588" data-script-author-name="gorgiaxx" data-script-daily-installs="307" data-script-total-installs="33801" data-script-rating-score="96.2" data-script-created-date="2018-09-22" data-script-updated-date="2019-02-14" data-script-type="public" data-script-version="1.2.9" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/372452-csdn%E8%87%AA%E5%8A%A8%E5%B1%95%E5%BC%80-%E5%8E%BB%E5%B9%BF%E5%91%8A-%E5%87%80%E5%8C%96%E5%89%AA%E8%B4%B4%E6%9D%BF-%E5%85%8D%E7%99%BB%E9%99%86">CSDN自动展开+去广告+净化剪贴板+免登陆</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				ITeye CSDN自动展开阅读，可以将剪贴板的推广信息去除，去除大多数广告。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/214588-gorgiaxx">gorgiaxx</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>307</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>33,801</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="96.2"><span><span class="good-rating-count" title="好评或收藏的人数。">144</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-09-22T06:52:43+00:00">2018-09-22</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-14T03:18:19+00:00">2019-02-14</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="26992" data-script-name="贴吧全能助手" data-script-author-id="99272" data-script-author-name="忆世萧遥" data-script-daily-installs="326" data-script-total-installs="211804" data-script-rating-score="98.7" data-script-created-date="2017-02-02" data-script-updated-date="2017-04-19" data-script-type="public" data-script-version="2.1" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/26992-%E8%B4%B4%E5%90%A7%E5%85%A8%E8%83%BD%E5%8A%A9%E6%89%8B">贴吧全能助手</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				【装这一个脚本就够了～可能是你遇到的最好用的贴吧增强脚本】，百度贴吧 tieba.baidu.com 看贴（包括楼中楼）无须登录，完全去除扰眼和各类广告模块，全面精简并美化各种贴吧页面，去除贴吧帖子里链接的跳转，按发帖时间排序，查看贴吧用户发言记录，贴子关键字屏蔽，移除会员彩名，直接在当前页面查看原图，可缩放，可多开，可拖拽
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/99272-%E5%BF%86%E4%B8%96%E8%90%A7%E9%81%A5">忆世萧遥</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>326</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>211,804</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.7"><span><span class="good-rating-count" title="好评或收藏的人数。">1293</span>
<span class="ok-rating-count" title="评级为一般的人数。">8</span>
<span class="bad-rating-count" title="评级为差评的人数。">5</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-02-02T03:27:56+00:00">2017-02-02</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2017-04-19T03:55:18+00:00">2017-04-19</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="373334" data-script-name="百度文库（wenku）在线下载PDF格式文件" data-script-author-id="218639" data-script-author-name="詹eko" data-script-daily-installs="304" data-script-total-installs="37957" data-script-rating-score="96.0" data-script-created-date="2018-10-17" data-script-updated-date="2019-01-11" data-script-type="public" data-script-version="0.1.7" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/373334-%E7%99%BE%E5%BA%A6%E6%96%87%E5%BA%93-wenku-%E5%9C%A8%E7%BA%BF%E4%B8%8B%E8%BD%BDpdf%E6%A0%BC%E5%BC%8F%E6%96%87%E4%BB%B6">百度文库（wenku）在线下载PDF格式文件</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				百度文库文档页面打印PDF，chrome浏览器最好能安装一下 adblock 插件，下载后的pdf文件可以在 https://pdf2docx.com/zh/ 上转换成docx
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/218639-%E8%A9%B9eko">詹eko</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>304</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>37,957</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="96.0"><span><span class="good-rating-count" title="好评或收藏的人数。">135</span>
<span class="ok-rating-count" title="评级为一般的人数。">2</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-10-17T02:00:17+00:00">2018-10-17</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-11T10:56:22+00:00">2019-01-11</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="18842" data-script-name="CNKI 中国知网 PDF 全文下载（特制版）" data-script-author-id="39034" data-script-author-name="Yue" data-script-daily-installs="292" data-script-total-installs="158568" data-script-rating-score="99.2" data-script-created-date="2016-04-17" data-script-updated-date="2017-02-10" data-script-type="public" data-script-version="3.2.0.20170210" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/18842-cnki-%E4%B8%AD%E5%9B%BD%E7%9F%A5%E7%BD%91-pdf-%E5%85%A8%E6%96%87%E4%B8%8B%E8%BD%BD-%E7%89%B9%E5%88%B6%E7%89%88">CNKI 中国知网 PDF 全文下载（特制版）</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				参见 http://blog.yuelong.info/post/cnki-pdf-js.html
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/39034-yue">Yue</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>292</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>158,568</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.2"><span><span class="good-rating-count" title="好评或收藏的人数。">960</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">2</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-04-17T03:00:42+00:00">2016-04-17</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2017-02-10T08:12:49+00:00">2017-02-10</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="14146" data-script-name="网页限制解除" data-script-author-id="20812" data-script-author-name="Cat73" data-script-daily-installs="298" data-script-total-installs="400818" data-script-rating-score="99.6" data-script-created-date="2015-11-23" data-script-updated-date="2016-12-18" data-script-type="public" data-script-version="1.3" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/14146-%E7%BD%91%E9%A1%B5%E9%99%90%E5%88%B6%E8%A7%A3%E9%99%A4">网页限制解除</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				通杀大部分网站，可以解除禁止复制、剪切、选择文本、右键菜单的限制。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/20812-cat73">Cat73</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>298</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>400,818</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.6"><span><span class="good-rating-count" title="好评或收藏的人数。">2054</span>
<span class="ok-rating-count" title="评级为一般的人数。">4</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2015-11-23T11:51:11+00:00">2015-11-23</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2016-12-18T20:08:05+00:00">2016-12-18</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="25270" data-script-name="城通网盘、皮皮盘、牛盘显示正确下载地址" data-script-author-id="21277" data-script-author-name="Dave B" data-script-daily-installs="263" data-script-total-installs="167693" data-script-rating-score="99.2" data-script-created-date="2016-11-30" data-script-updated-date="2018-07-18" data-script-type="public" data-script-version="0.9.1" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/25270-%E5%9F%8E%E9%80%9A%E7%BD%91%E7%9B%98-%E7%9A%AE%E7%9A%AE%E7%9B%98-%E7%89%9B%E7%9B%98%E6%98%BE%E7%A4%BA%E6%AD%A3%E7%A1%AE%E4%B8%8B%E8%BD%BD%E5%9C%B0%E5%9D%80">城通网盘、皮皮盘、牛盘显示正确下载地址</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				城通网盘、皮皮盘、牛盘显示正确下载地址!去掉遮挡的popjump透明层！自动跳到第二页！
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/21277-dave-b">Dave B</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>263</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>167,693</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.2"><span><span class="good-rating-count" title="好评或收藏的人数。">1040</span>
<span class="ok-rating-count" title="评级为一般的人数。">3</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-11-30T06:48:00+00:00">2016-11-30</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-07-18T16:00:09+00:00">2018-07-18</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="18991" data-script-name="[MTurk Worker] Dashboard Enhancer" data-script-author-id="11580" data-script-author-name="Kadauchi" data-script-daily-installs="263" data-script-total-installs="27726" data-script-rating-score="76.4" data-script-created-date="2016-04-21" data-script-updated-date="2019-02-19" data-script-type="public" data-script-version="2.0.5" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/18991-mturk-worker-dashboard-enhancer">[MTurk Worker] Dashboard Enhancer</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Brings many enhancements to the MTurk Worker Dashboard.
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/11580-kadauchi">Kadauchi</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>263</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>27,726</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="76.4"><span><span class="good-rating-count" title="好评或收藏的人数。">19</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-04-21T03:49:39+00:00">2016-04-21</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-19T22:30:41+00:00">2019-02-19</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="375075" data-script-name="SURVIV.IO MODS / HACKS / CHEATS by MR.Coder | FAST HEAL, AIM LOCK, FIRE BOT, AUTO RESPAWN" data-script-author-id="228916" data-script-author-name="Mr.Coder" data-script-daily-installs="276" data-script-total-installs="23825" data-script-rating-score="40.2" data-script-created-date="2018-12-01" data-script-updated-date="2019-01-27" data-script-type="public" data-script-version="2.2.2" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/375075-surviv-io-mods-hacks-cheats-by-mr-coder-fast-heal-aim-lock-fire-bot-auto-respawn">SURVIV.IO MODS / HACKS / CHEATS by MR.Coder | FAST HEAL, AIM LOCK, FIRE BOT, AUTO RESPAWN</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Survivio Hacks for surviv.io cheats with surviv mods that aim lock,fire bot, auto respawn, auto reload, show fps, fast heal, auto weapon swap
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/228916-mr-coder">Mr.Coder</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>276</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>23,825</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="40.2"><span><span class="good-rating-count" title="好评或收藏的人数。">6</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">2</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-12-01T23:25:08+00:00">2018-12-01</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-27T02:49:54+00:00">2019-01-27</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="372516" data-script-name="bilibili merged flv+mp4+ass+enhance" data-script-author-id="215003" data-script-author-name="Xmader" data-script-daily-installs="253" data-script-total-installs="25615" data-script-rating-score="96.9" data-script-created-date="2018-09-24" data-script-updated-date="2019-02-15" data-script-type="public" data-script-version="1.21.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/372516-bilibili-merged-flv-mp4-ass-enhance">bilibili merged flv+mp4+ass+enhance</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				bilibili/哔哩哔哩:超清FLV下载,FLV合并,原生MP4下载,弹幕ASS下载,MKV打包,播放体验增强,原生appsecret,不借助其他网站
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/215003-xmader">Xmader</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>253</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>25,615</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="96.9"><span><span class="good-rating-count" title="好评或收藏的人数。">122</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-09-24T08:43:51+00:00">2018-09-24</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-15T01:54:15+00:00">2019-02-15</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="39776" data-script-name="百度网盘直接下载助手修改版" data-script-author-id="145237" data-script-author-name="yi ming" data-script-daily-installs="243" data-script-total-installs="427077" data-script-rating-score="98.1" data-script-created-date="2018-03-20" data-script-updated-date="2018-04-06" data-script-type="public" data-script-version="2018.04.06" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/39776-%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E7%9B%B4%E6%8E%A5%E4%B8%8B%E8%BD%BD%E5%8A%A9%E6%89%8B%E4%BF%AE%E6%94%B9%E7%89%88">百度网盘直接下载助手修改版</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				网盘内和分享页均显示[下载助手]按钮，支持获取直接下载链接+压缩下载链接；
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/145237-yi-ming">yi ming</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>243</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>427,077</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="98.1"><span><span class="good-rating-count" title="好评或收藏的人数。">695</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">6</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-03-20T08:53:14+00:00">2018-03-20</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-04-06T06:09:15+00:00">2018-04-06</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="376724" data-script-name="VIP视频免费看（全网）-by无心-【2019-02-21更新可用】" data-script-author-id="221175" data-script-author-name="我本无心" data-script-daily-installs="241" data-script-total-installs="7871" data-script-rating-score="83.9" data-script-created-date="2019-01-15" data-script-updated-date="2019-02-21" data-script-type="public" data-script-version="20.19.02.26" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/376724-vip%E8%A7%86%E9%A2%91%E5%85%8D%E8%B4%B9%E7%9C%8B-%E5%85%A8%E7%BD%91-by%E6%97%A0%E5%BF%83-2019-02-21%E6%9B%B4%E6%96%B0%E5%8F%AF%E7%94%A8">VIP视频免费看（全网）-by无心-【2019-02-21更新可用】</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				【待添加多功能备用链】【新增插件下载】【新增接口大全】【新增音乐解析】【新增悬浮球】【多线解析】在视频标题旁上红色字体显示网站解析、接口解析、直播解析、音乐解析、我要反馈等按钮.采用网站一站式解析和多接口集成的方式直接调用接口，接口全部亲自测试过，同时接口支持长期的更新替换工作。在保证接口解析质量的前提下，不断地优化界面以及提高用户使用体验。因为只专注于vip影视解析，所以更专业更值得信赖。注：接口全部来自该网站：http://www.luckyblank.cn/vip_videos。Watch VIP movies for free, using website parsing and interface parsing. Is it more trustworthy because it focuses on parsing?
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/221175-%E6%88%91%E6%9C%AC%E6%97%A0%E5%BF%83">我本无心</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>241</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>7,871</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="83.9"><span><span class="good-rating-count" title="好评或收藏的人数。">20</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2019-01-15T08:27:09+00:00">2019-01-15</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-21T03:52:20+00:00">2019-02-21</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="377137" data-script-name="Skribbl.io Draw Bot, Auto Guess, Word Helper 2019 (Mods,Hacks,Cheats) by SKRIBBL-IO.NET" data-script-author-id="228916" data-script-author-name="Mr.Coder" data-script-daily-installs="241" data-script-total-installs="5981" data-script-rating-score="5.0" data-script-created-date="2019-01-26" data-script-updated-date="2019-01-27" data-script-type="public" data-script-version="1.2" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/377137-skribbl-io-draw-bot-auto-guess-word-helper-2019-mods-hacks-cheats-by-skribbl-io-net">Skribbl.io Draw Bot, Auto Guess, Word Helper 2019 (Mods,Hacks,Cheats) by SKRIBBL-IO.NET</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Skribblio Cheats, Hacks, Mods, Unblocked 2019 Game Play Skribbl.io Draw Bot, Auto Guesser, Word Helper
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/228916-mr-coder">Mr.Coder</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>241</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>5,981</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="5.0"><span><span class="good-rating-count" title="好评或收藏的人数。">0</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2019-01-26T04:23:27+00:00">2019-01-26</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-27T05:23:29+00:00">2019-01-27</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="370308" data-script-name="QQ音乐付费无损音乐免费下载" data-script-author-id="165079" data-script-author-name="刘明野" data-script-daily-installs="269" data-script-total-installs="24391" data-script-rating-score="93.5" data-script-created-date="2018-07-15" data-script-updated-date="2018-09-19" data-script-type="public" data-script-version="1.0.5" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/370308-qq%E9%9F%B3%E4%B9%90%E4%BB%98%E8%B4%B9%E6%97%A0%E6%8D%9F%E9%9F%B3%E4%B9%90%E5%85%8D%E8%B4%B9%E4%B8%8B%E8%BD%BD">QQ音乐付费无损音乐免费下载</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				提供QQ音乐付费歌曲，320K下载、无损FLAC下载、无损APE下载、歌词下载，可以解析歌单列表、歌手歌单、分类歌单、专辑列表
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/165079-%E5%88%98%E6%98%8E%E9%87%8E">刘明野</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>269</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>24,391</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="93.5"><span><span class="good-rating-count" title="好评或收藏的人数。">115</span>
<span class="ok-rating-count" title="评级为一般的人数。">3</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-07-15T15:16:13+00:00">2018-07-15</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-09-19T08:08:01+00:00">2018-09-19</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="370811" data-script-name="【2019年2月19日更新】网盘万能钥匙，自动查询百度网盘分享链接的提取码,全网VIP视频解析播放,全网付费音乐免费下载,淘宝、拼多多大额购物优惠券领取，支持历史价格查询" data-script-author-id="194233" data-script-author-name="一起购物" data-script-daily-installs="216" data-script-total-installs="46265" data-script-rating-score="89.9" data-script-created-date="2018-08-01" data-script-updated-date="2019-02-18" data-script-type="public" data-script-version="3.1.5" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/370811-2019%E5%B9%B42%E6%9C%8819%E6%97%A5%E6%9B%B4%E6%96%B0-%E7%BD%91%E7%9B%98%E4%B8%87%E8%83%BD%E9%92%A5%E5%8C%99-%E8%87%AA%E5%8A%A8%E6%9F%A5%E8%AF%A2%E7%99%BE%E5%BA%A6%E7%BD%91%E7%9B%98%E5%88%86%E4%BA%AB%E9%93%BE%E6%8E%A5%E7%9A%84%E6%8F%90%E5%8F%96%E7%A0%81-%E5%85%A8%E7%BD%91vip%E8%A7%86%E9%A2%91%E8%A7%A3%E6%9E%90%E6%92%AD%E6%94%BE-%E5%85%A8%E7%BD%91%E4%BB%98%E8%B4%B9%E9%9F%B3%E4%B9%90%E5%85%8D%E8%B4%B9%E4%B8%8B%E8%BD%BD-%E6%B7%98%E5%AE%9D-%E6%8B%BC%E5%A4%9A%E5%A4%9A%E5%A4%A7%E9%A2%9D%E8%B4%AD%E7%89%A9%E4%BC%98%E6%83%A0%E5%88%B8%E9%A2%86%E5%8F%96-%E6%94%AF%E6%8C%81%E5%8E%86%E5%8F%B2%E4%BB%B7%E6%A0%BC%E6%9F%A5%E8%AF%A2">【2019年2月19日更新】网盘万能钥匙，自动查询百度网盘分享链接的提取码,全网VIP视频解析播放,全网付费音乐免费下载,淘宝、拼多多大额购物优惠券领取，支持历史价格查询</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				自动查询百度网盘分享链接的提取码,是网盘界的万能钥匙,全网VIP视频解析播放,全网付费音乐免费下载,淘宝、拼多多大额购物优惠券领取，支持商品比价,查询历史价格查询,使用优惠券购物,脚本开发者会获得些许返利.支持油猴、暴力猴插件。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/194233-%E4%B8%80%E8%B5%B7%E8%B4%AD%E7%89%A9">一起购物</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>216</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>46,265</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="89.9"><span><span class="good-rating-count" title="好评或收藏的人数。">80</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">3</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-08-01T14:10:15+00:00">2018-08-01</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-18T11:24:53+00:00">2019-02-18</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="30578" data-script-name="解决百度云大文件下载限制 （路人添加版）" data-script-author-id="132323" data-script-author-name="cotsnow" data-script-daily-installs="224" data-script-total-installs="132499" data-script-rating-score="97.0" data-script-created-date="2017-06-13" data-script-updated-date="2017-06-15" data-script-type="public" data-script-version="0.0.6+" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/30578-%E8%A7%A3%E5%86%B3%E7%99%BE%E5%BA%A6%E4%BA%91%E5%A4%A7%E6%96%87%E4%BB%B6%E4%B8%8B%E8%BD%BD%E9%99%90%E5%88%B6-%E8%B7%AF%E4%BA%BA%E6%B7%BB%E5%8A%A0%E7%89%88">解决百度云大文件下载限制 （路人添加版）</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				一行代码，解决百度云大文件下载限制
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/132323-cotsnow">cotsnow</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>224</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>132,499</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="97.0"><span><span class="good-rating-count" title="好评或收藏的人数。">123</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-06-13T14:13:18+00:00">2017-06-13</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2017-06-15T09:51:29+00:00">2017-06-15</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="21958" data-script-name="Split dan Macro untuk Bubble.am" data-script-author-id="58428" data-script-author-name="TukangTIDUR" data-script-daily-installs="224" data-script-total-installs="231621" data-script-rating-score="90.6" data-script-created-date="2016-08-03" data-script-updated-date="2016-08-03" data-script-type="public" data-script-version="2" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/21958-split-dan-macro-untuk-bubble-am">Split dan Macro untuk Bubble.am</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Shift is split, Q is macro
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/58428-tukangtidur">TukangTIDUR</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>224</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>231,621</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="90.6"><span><span class="good-rating-count" title="好评或收藏的人数。">55</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-08-03T16:49:42+00:00">2016-08-03</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2016-08-03T16:51:36+00:00">2016-08-03</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="37654" data-script-name="弹窗阻止程序脚本" data-script-author-id="167601" data-script-author-name="MikeWang87" data-script-daily-installs="228" data-script-total-installs="116837" data-script-rating-score="97.8" data-script-created-date="2018-01-21" data-script-updated-date="2018-02-26" data-script-type="public" data-script-version="0.2" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/37654-popup-blocker-script">弹窗阻止程序脚本</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				用于阻止所有类型弹窗的最有效用户脚本。为防范最狡猾的弹窗而设计，包括成人和流媒体网站上的弹窗。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/167601-mikewang87">MikeWang87</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>228</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>116,837</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="97.8"><span><span class="good-rating-count" title="好评或收藏的人数。">498</span>
<span class="ok-rating-count" title="评级为一般的人数。">3</span>
<span class="bad-rating-count" title="评级为差评的人数。">3</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-01-21T10:27:11+00:00">2018-01-21</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-02-26T10:33:12+00:00">2018-02-26</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="376482" data-script-name="Krunker.io 2019 Hacks, Cheats, Mods | NO DISCONNECT ISSUE | AIMBOT KRUNKERIO.NET" data-script-author-id="228916" data-script-author-name="Mr.Coder" data-script-daily-installs="196" data-script-total-installs="5222" data-script-rating-score="4.6" data-script-created-date="2019-01-08" data-script-updated-date="2019-02-04" data-script-type="public" data-script-version="5.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/376482-krunker-io-2019-hacks-cheats-mods-no-disconnect-issue-aimbot-krunkerio-net">Krunker.io 2019 Hacks, Cheats, Mods | NO DISCONNECT ISSUE | AIMBOT KRUNKERIO.NET</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Fire Bot, Aim Bot, Unlimited Ammo, ESP, Wall Hack, Auto Heal... new working 2019 krunkerio hack with adblock
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/228916-mr-coder">Mr.Coder</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>196</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>5,222</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="4.6"><span><span class="good-rating-count" title="好评或收藏的人数。">1</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">3</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2019-01-08T14:59:39+00:00">2019-01-08</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-04T07:49:05+00:00">2019-02-04</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="375240" data-script-name="移除百家号搜索结果" data-script-author-id="230778" data-script-author-name="mzcc" data-script-daily-installs="224" data-script-total-installs="12112" data-script-rating-score="92.3" data-script-created-date="2018-12-06" data-script-updated-date="2019-02-20" data-script-type="public" data-script-version="1.10" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/375240-%E7%A7%BB%E9%99%A4%E7%99%BE%E5%AE%B6%E5%8F%B7%E6%90%9C%E7%B4%A2%E7%BB%93%E6%9E%9C">移除百家号搜索结果</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				删除百度搜索结果的百家号结果
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/230778-mzcc">mzcc</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>224</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>12,112</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="92.3"><span><span class="good-rating-count" title="好评或收藏的人数。">57</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">0</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-12-06T05:27:06+00:00">2018-12-06</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-20T02:40:09+00:00">2019-02-20</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="373063" data-script-name="视频跳过广告和 VIP 视频解析" data-script-author-id="197659" data-script-author-name="mofiter" data-script-daily-installs="172" data-script-total-installs="37509" data-script-rating-score="92.6" data-script-created-date="2018-10-09" data-script-updated-date="2019-01-16" data-script-type="public" data-script-version="1.3" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/373063-%E8%A7%86%E9%A2%91%E8%B7%B3%E8%BF%87%E5%B9%BF%E5%91%8A%E5%92%8C-vip-%E8%A7%86%E9%A2%91%E8%A7%A3%E6%9E%90">视频跳过广告和 VIP 视频解析</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				也许是风格最贴近原视频网站的 VIP 视频解析脚本了，添加的解析按钮样式跟原视频网站已有的按钮一致，不会产生突兀感，支持腾讯视频、爱奇艺、优酷、土豆、芒果 TV、搜狐视频、乐视视频、PPTV等，支持多个解析接口切换，支持自定义接口，支持站内站外解析，支持 Tampermonkey、Violentmonkey、Greasemonkey
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/197659-mofiter">mofiter</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>172</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>37,509</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="92.6"><span><span class="good-rating-count" title="好评或收藏的人数。">82</span>
<span class="ok-rating-count" title="评级为一般的人数。">1</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-10-09T14:59:51+00:00">2018-10-09</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-16T05:25:16+00:00">2019-01-16</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="371356" data-script-name="BEST moomoo.io mod/hack/cheat! Spike/trap hotkey, pro map, autoheal, autobull/monkey tail! (2018)" data-script-author-id="206156" data-script-author-name="Long Nhat" data-script-daily-installs="185" data-script-total-installs="12166" data-script-rating-score="20.8" data-script-created-date="2018-08-20" data-script-updated-date="2018-08-20" data-script-type="public" data-script-version="10.0" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/371356-best-moomoo-io-mod-hack-cheat-spike-trap-hotkey-pro-map-autoheal-autobull-monkey-tail-2018">BEST moomoo.io mod/hack/cheat! Spike/trap hotkey, pro map, autoheal, autobull/monkey tail! (2018)</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				Autoheal
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/206156-long-nhat">Long Nhat</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>185</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>12,166</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="20.8"><span><span class="good-rating-count" title="好评或收藏的人数。">2</span>
<span class="ok-rating-count" title="评级为一般的人数。">0</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2018-08-20T03:11:44+00:00">2018-08-20</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-08-20T03:11:44+00:00">2018-08-20</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="24508" data-script-name="Userscript+ : 显示当前网站所有可用的UserJS脚本 Jaeger" data-script-author-id="7267" data-script-author-name="jaeger" data-script-daily-installs="170" data-script-total-installs="104321" data-script-rating-score="99.4" data-script-created-date="2016-11-02" data-script-updated-date="2018-10-07" data-script-type="public" data-script-version="2.3.3" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/24508-userscript-show-site-all-userjs">Userscript+ : 显示当前网站所有可用的UserJS脚本 Jaeger</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				显示当前网站的所有可用UserJS(Tampermonkey)脚本,交流QQ群:104267383
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/7267-jaeger">jaeger</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>170</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>104,321</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.4"><span><span class="good-rating-count" title="好评或收藏的人数。">1300</span>
<span class="ok-rating-count" title="评级为一般的人数。">2</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2016-11-02T06:51:53+00:00">2016-11-02</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2018-10-07T14:35:06+00:00">2018-10-07</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="3249" data-script-name="Yet Another Weibo Filter 看真正想看的微博" data-script-author-id="198" data-script-author-name="ts" data-script-daily-installs="170" data-script-total-installs="104828" data-script-rating-score="99.3" data-script-created-date="2014-07-12" data-script-updated-date="2019-01-31" data-script-type="public" data-script-version="3.7.493" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/3249-yet-another-weibo-filter">Yet Another Weibo Filter 看真正想看的微博</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				新浪微博根据关键词、作者、话题、来源等过滤微博；修改版面。
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/198-ts">ts</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>170</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>104,828</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="99.3"><span><span class="good-rating-count" title="好评或收藏的人数。">1095</span>
<span class="ok-rating-count" title="评级为一般的人数。">2</span>
<span class="bad-rating-count" title="评级为差评的人数。">1</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2014-07-12T12:19:00+00:00">2014-07-12</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-01-31T12:30:25+00:00">2019-01-31</time></span></dd>
		</dl>
	</article>
</li>
<li data-script-id="34613" data-script-name="YouTube Multi Downloader v4.4 (MP3, FHD, AVI, MP4, FLV, WKV, ACC, M4V, etc) - Saveclipbro.com" data-script-author-id="152924" data-script-author-name="Punisher" data-script-daily-installs="172" data-script-total-installs="52204" data-script-rating-score="88.9" data-script-created-date="2017-10-29" data-script-updated-date="2019-02-06" data-script-type="public" data-script-version="4.4" data-sensitive="false">
	<article>
		<h2>
			<a href="/zh-CN/scripts/34613-youtube-multi-downloader-v4-4-mp3-fhd-avi-mp4-flv-wkv-acc-m4v-etc-saveclipbro-com">YouTube Multi Downloader v4.4 (MP3, FHD, AVI, MP4, FLV, WKV, ACC, M4V, etc) - Saveclipbro.com</a>
			<span class="name-description-separator">
				-
			</span>
			<span class="description">
				No ads, this script helps to add a download button. Saveclipbro.com features the fastest ways to download YouTube videos and audios, ensuring quality. Supported Services: YouTube, Dailymotion, Vimeo, Twitter, Vine, Facebook, Instagram, Xvideos, CNN, SoundCloud, Redtube, Freesound, Pornhub, Reddit, Liveleak and more! You will like!
			</span>
		</h2>
		<dl class="inline-script-stats">
			<dt class="script-list-author"><span>作者</span></dt>
			<dd class="script-list-author"><span><a href="/zh-CN/users/152924-punisher">Punisher</a></span></dd>
				<dt class="script-list-daily-installs"><span>今日安装</span></dt>
				<dd class="script-list-daily-installs"><span>172</span></dd>
				<dt class="script-list-total-installs"><span>总安装量</span></dt>
				<dd class="script-list-total-installs"><span>52,204</span></dd>
				<dt class="script-list-ratings"><span>得分</span></dt>
				<dd class="script-list-ratings" data-rating-score="88.9"><span><span class="good-rating-count" title="好评或收藏的人数。">117</span>
<span class="ok-rating-count" title="评级为一般的人数。">2</span>
<span class="bad-rating-count" title="评级为差评的人数。">6</span>
</span></dd>
			<dt class="script-list-created-date"><span>创建日期</span></dt>
			<dd class="script-list-created-date"><span><time datetime="2017-10-29T00:18:04+00:00">2017-10-29</time></span></dd>
			<dt class="script-list-updated-date"><span>最近更新</span></dt>
			<dd class="script-list-updated-date"><span><time datetime="2019-02-06T02:01:57+00:00">2019-02-06</time></span></dd>
		</dl>
	</article>
</li>

      </ol>

      <div class="pagination"><span class="previous_page disabled">上一页</span> <em class="current">1</em> <a rel="next" href="/zh-CN/scripts?page=2">2</a> <a href="/zh-CN/scripts?page=3">3</a> <a href="/zh-CN/scripts?page=4">4</a> <a href="/zh-CN/scripts?page=5">5</a> <a href="/zh-CN/scripts?page=6">6</a> <a href="/zh-CN/scripts?page=7">7</a> <a href="/zh-CN/scripts?page=8">8</a> <a href="/zh-CN/scripts?page=9">9</a> <span class="gap">…</span> <a href="/zh-CN/scripts?page=19">19</a> <a href="/zh-CN/scripts?page=20">20</a> <a class="next_page" rel="next" href="/zh-CN/scripts?page=2">下一页</a></div>

    <p><a href="/zh-CN/script_versions/new">发布您编写的脚本</a>(或 <a href="/zh-CN/help/writing-user-scripts">了解如何编写脚本</a>)</p>
  </div>
  <div class="sidebar collapsed">
    <div class="close-sidebar">
      <div class="sidebar-title">搜索选项</div>
      <div>☰</div>
    </div>
    
<div id="script-list-option-groups" class="list-option-groups">

		<form class="sidebar-search">
				
				
			<input type="search" name="q" value="" placeholder="搜索脚本"><input class="search-submit" type="submit" value="🔎">
		</form>

	<div id="script-list-sort" class="list-option-group">脚本排序：
		<ul>
				<li class="list-option list-current">今日安装</li>
				<li class="list-option"><a href="/zh-CN/scripts?sort=total_installs">总安装数</a></li>
				<li class="list-option"><a href="/zh-CN/scripts?sort=ratings">得分</a></li>
				<li class="list-option"><a href="/zh-CN/scripts?sort=created">创建日期</a></li>
				<li class="list-option"><a href="/zh-CN/scripts?sort=updated">更新日期</a></li>
				<li class="list-option"><a href="/zh-CN/scripts?sort=name">名称</a></li>
		</ul>
	</div>

		<div class="ad carbon-ad">
			<script async type="text/javascript" src="//cdn.carbonads.com/carbon.js?serve=CKYI5KQU&placement=greasyforkorg" id="_carbonads_js"></script>
		</div>

		<div id="script-list-filter" class="list-option-group">根据脚本生效的网站过滤：
			<ul>
				<li class="list-option list-current">全部</li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/baidu.com">baidu.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/qq.com">qq.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/bilibili.com">bilibili.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/sohu.com">sohu.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/youku.com">youku.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/le.com">le.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/mgtv.com">mgtv.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/tudou.com">tudou.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/iqiyi.com">iqiyi.com</a></li>
						<li class="list-option"><a href="/zh-CN/scripts/by-site/pptv.com">pptv.com</a></li>
				<li><a href="/zh-CN/scripts/by-site">更多...</a></li>
			</ul>
		</div>

</div>


  </div>
</div>


  </div>

    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

      ga('create', 'UA-48197018-1', 'greasyfork.org');
      ga('send', 'pageview');

    </script>

</body>
</html>

`//<option></option>
	//reg2:=regexp.MustCompile(`<option>.</option>`)
	reg2:=regexp.MustCompile(`</option>(?s:(.*?))</option>`)
	if reg2 ==nil{
		fmt.Println("error")
		return
	}
	s3:=reg2.FindAllStringSubmatch(tmp3,-1)
	//fmt.Println("s3=",s3)
	for _,text:=range s3{
		fmt.Println(text[1])//实现内容筛选，去掉</option>
	}
}
