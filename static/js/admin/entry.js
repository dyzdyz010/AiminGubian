$(document).ready(function () {
	$('.datepicker').pickadate({
		format: 'yyyy-mm-dd',
    	selectMonths: true, // Creates a dropdown to control month
    	selectYears: 15 // Creates a dropdown of 15 years to control year
    });


	// Image upload
	imgUploadInit();
	console.log($('#submit-btn'));

	$('form').submit(function (e) {
		if (!$('#title').val()) {
			alert('标题不能为空。');
			e.preventDefault();
		};
	});
});