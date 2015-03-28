<!doctype html>
<html>
<head>
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />

	<!-- Style Sheets -->
	{{.Stylesheets}}
</head>
<body>
	<nav>
		<div class="nav-wrapper container">
			<a href="/" class="brand-logo">中国边防</a>
			<a href="#" data-activates="mobile" class="button-collapse"><i class="mdi-navigation-menu"></i></a>
			<ul id="nav-mobile" class="right hide-on-med-and-down">
				<li><a href="/police-station">派出所简介</a></li>
				<li><a href="/express">直通车</a></li>
				<li><a href="/express/timeline">时间线</a></li>
			</ul>
			<ul class="side-nav" id="mobile">
				<li><a href="/police-station">派出所简介</a></li>
				<li><a href="/express">直通车</a></li>
				<li><a href="/express/timeline">时间线</a></li>
			</ul>
		</div>
	</nav>

	{{.LayoutContent}}

	<footer class="page-footer">
		<div class="container">
			<div class="row">
				<div class="col l6 s12">
					<h5 class="white-text">帮助网站进步</h5>
					<p class="grey-text text-lighten-4">欢迎广大读者将此网站分享给更多的人。</p>
					<img width="100px" src="https://chart.googleapis.com/chart?cht=qr&chl=&chs=180x180&choe=UTF-8&chld=L|2">
				</div>
				<div class="col l4 offset-l2 s12">
					<h5 class="white-text">关于作者</h5>
					<ul>
						<li><a class="grey-text text-lighten-3" href="http://blog.duyizhuo.com">杜艺卓</a></li>
						<li><a class="grey-text text-lighten-3" href="#!">Github</a></li>
					</ul>
				</div>
			</div>
		</div>
		<div class="footer-copyright">
			<div class="container">
				© 2015 Copyright 巴彦淖尔边防支队
			</div>
		</div>
	</footer>

	{{.Scripts}}
</body>