function formValueTest() {
	let dataList = [];
	const obj1 = {name: "メイクデビュー", result: 1};
	dataList.push(obj1);
	const obj2 = {name: "ホープフルステークス", result: 1};
	dataList.push(obj2);
	
	const e = document.getElementById("hello_page_list").value = JSON.stringify(dataList);
}
