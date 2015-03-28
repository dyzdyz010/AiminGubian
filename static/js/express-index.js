$(document).ready(function () {
	$('.scrollspy').scrollSpy();
	$('.pinmenu .table-of-contents').pushpin({ top: $('.pinmenu').offset().top });
	$('.modal-trigger').leanModal();
});