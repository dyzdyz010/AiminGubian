$(document).ready(function () {
	// Prepare layout options.
	var wookmark;

	// Init lightbox
	$('#gallery-container').magnificPopup({
		delegate: 'li:not(.inactive) a',
		type: 'image',
		gallery: {
			enabled: true
		}
	});

	// Call the layout function after all images have loaded
	imagesLoaded('#gallery-container', function () {
		wookmark = new Wookmark('#gallery-container', {
		offset: 2, // Optional, the distance between grid items
		itemWidth: 210 // Optional, the width of a grid item
	});
	});
});