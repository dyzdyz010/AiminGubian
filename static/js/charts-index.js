$(document).ready(function () {
	$('.scrollspy').scrollSpy();
	$('.pinmenu .table-of-contents').pushpin({ top: $('.pinmenu').offset().top });

	var date = new Date();
	var month = date.getMonth() - 1;

	var purchaseDataSource = [
	{month: "一月",   food: 3, staff: 1},
	{month: "二月",   food: 2, staff: 12},
	{month: "三月",   food: 3, staff: 14},
	{month: "四月",   food: 4, staff: 12},
	{month: "五月",   food: 6, staff: 20},
	{month: "六月",   food: 11, staff: 15},
	{month: "七月",   food: 4, staff: 11},
	{month: "八月",   food: 7, staff: 12},
	{month: "九月",   food: 10, staff: 11},
	{month: "十月",   food: 1, staff: 10},
	{month: "十一月", food: 29, staff: 5},
	{month: "十二月", food: 20, staff: 4}];

	$('#chartProfitContainer').dxChart({
		dataSource: purchaseDataSource,
		commonSeriesSettings: {
			argumentField: "month",
			type: "stackedBar"
		},

		series: [
		{valueField: "food", name: "食品", color: '#ffa500'},
		{valueField: "staff", name: "生活用品", color: '#34b3f4'}],

		title: "近一年直通车运行购买金额（元）"
	});

	var costDataSource = [
	{month: "一月",   oil: 40, food: 70},
	{month: "二月",   oil: 80, food: 70},
	{month: "三月",   oil: 76, food: 70},
	{month: "四月",   oil: 35, food: 70},
	{month: "五月",   oil: 46, food: 70},
	{month: "六月",   oil: 45, food: 70},
	{month: "七月",   oil: 75, food: 70},
	{month: "八月",   oil: 31, food: 70},
	{month: "九月",   oil: 55, food: 70},
	{month: "十月",   oil: 77, food: 70},
	{month: "十一月", oil: 40, food: 70},
	{month: "十二月", oil: 90, food: 70}];

	$('#chartCostContainer').dxChart({
		dataSource: costDataSource,
		commonSeriesSettings: {
			argumentField: "month",
			type: "stackedBar"
		},

		series: [{
			valueField: "food", 
			name: "餐饮费", 
			color: '#ffa500'
		},{
			valueField: "oil", 
			name: "油料费", 
			color: "#34b3f4"
		}],

		title: "近一年直通车运行成本（元）"
	});

	var distanceDataSource = [
	{month: "一月",   distance: 80, },
	{month: "二月",   distance: 65, },
	{month: "三月",   distance: 34, },
	{month: "四月",   distance: 99, },
	{month: "五月",   distance: 37, },
	{month: "六月",   distance: 19, },
	{month: "七月",   distance: 100,},
	{month: "八月",   distance: 65, },
	{month: "九月",   distance: 76, },
	{month: "十月",   distance: 79, },
	{month: "十一月", distance: 64, },
	{month: "十二月", distance: 50, }];

	$('#chartDistanceContainer').dxChart({
		dataSource: distanceDataSource,
		commonSeriesSettings: {
			argumentField: "month",
			type: "bar"
		},

		series: {
			valueField: "distance", 
			name: "行驶里程(km)", 
			color: '#ffa500'
		},

		title: "近一年直通车行驶里程数"
	});
});