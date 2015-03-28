$(document).ready(function () {
	$('.scrollspy').scrollSpy();
	$('.pinmenu .table-of-contents').pushpin({ top: $('.pinmenu').offset().top });

	var map = new BMap.Map("placeMap");    // 创建Map实例
	map.setCurrentCity("巴彦淖尔");          // 设置地图显示的城市 此项是必须设置的
	var mapType = new BMap.MapTypeControl({mapTypes: [BMAP_NORMAL_MAP,BMAP_HYBRID_MAP]});
	map.addControl(mapType)
	var point = new BMap.Point(107.533671, 41.81794);
	map.centerAndZoom(point, 7);
	var myIcon = new BMap.Icon("/static/img/map-marker.png", new BMap.Size(30, 30), {    
		// 指定定位位置。   
		// 当标注显示在地图上时，其所指向的地理位置距离图标左上    
		// 角各偏移10像素和25像素。您可以看到在本例中该位置即是   
		// 图标中央下端的尖角位置。    
		offset: new BMap.Size(15, 30),    
		// 设置图片偏移。   
		// 当您需要从一幅较大的图片中截取某部分作为标注图标时，您   
		// 需要指定大图的偏移位置，此做法与css sprites技术类似。    
		// imageOffset: new BMap.Size(0, 0 - index * 25)   // 设置图片偏移    
	});
	var marker = new BMap.Marker(point, {icon: myIcon});// 创建标注
	map.addOverlay(marker);             // 将标注添加到地图中
	marker.disableDragging();           // 不可拖拽
});