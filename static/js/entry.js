$(document).ready(function () {
	var markedStr = marked($('.entry-content').text());
	var markDom = $(markedStr);
	$('.entry-content').html(markDom);
});