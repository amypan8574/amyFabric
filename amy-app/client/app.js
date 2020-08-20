// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller 頁面處理
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();

	$scope.urlFormat = function (name) {
		$scope.urlname = encodeURI(name)
	   } 
	
	$scope.queryAllCardPayment = function(){

		var channelID = $scope.queryCPR_channel;
		console.log(channelID);
        
		appFactory.queryAllCardPayment(channelID,function(data){
			var array = [];
			
			console.log(data);	
			//data.replaceAll(" ", "");
			//data = data.trim();
			//var testa = '[{"Key":"CPR1user1", "Record":{"acquiringBank":"Abank","acquiringDate":"20200407","ipfsHash":"ddfffff","objectType":"CPR","size":10000,"userId":"user1"}}]';

           data = JSON.parse(data);
		   console.log(typeof(data));
			for (var i = 0; i < data.length; i++){
				//parseInt(data[i].Key);
				data[i].Record.Key = data[i].Key;
				array.push(data[i].Record);
				console.log(data[i].Record);
			}
			// array.sort(function(a, b) {
			//     return parseFloat(a.Key) - parseFloat(b.Key);
			// });
			$scope.all_cardpayment = array;
		});
	}

	$scope.queryAllSR = function(){
        
		appFactory.queryAllSR(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				//parseInt(data[i].Key);
				data[i].Record.Key = data[i].Key;
				array.push(data[i].Record);
				
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_settlementreport = array;
		});
	}

	$scope.queryAllStoreData = function(){
        
		appFactory.queryAllStoreData(function(data){
			data = JSON.parse(data);
			var array = [];
			for (var i = 0; i < data.length; i++){
				//parseInt(data[i].Key);
				data[i].Record.Key = data[i].Key;
				array.push(data[i].Record);
				
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_storedata = array;
		});
	}

	$scope.queryCardPayment = function(){

		var id = $scope.cardpayment_id;

		appFactory.queryCardPayment(id, function(data){
			$scope.query_cardpayment = data;

			if ($scope.query_cardpayment == "Could not locate tuna"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordCardPayment = function(){
		$scope.file = document.getElementById('CPRfile').files[0];
		appFactory.recordCardPayment($scope.cardpayment,$scope.file, function(data){
			 $scope.create_cardpayment= data;
			$("#success_create").show();
		});
	}

	$scope.recordSettlementReport = function(){
		$scope.file = document.getElementById('SRfile').files[0];
		appFactory.recordSettlementReport($scope.settlementreport,$scope.file, function(data){
			 $scope.create_settlementreport= data;
			$("#success_create").show();
		});
	}

	$scope.recordStoreData = function(){
		$scope.file = document.getElementById('SDfile').files[0];
		appFactory.recordStoreData($scope.storedata,$scope.file, function(data){
			 //$scope.create_settlementreport= data;
			$("#success_create").show();
		});
	}

	$scope.changeValue = function(){

		appFactory.changeValue($scope.value, function(data){
			$scope.change_value = data;
			if ($scope.change_value == "Error: no tuna catch found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}
	$scope.add = function(file) {

		var fd = new FormData();
		fd.append('file', file);
	
		$http.post('/add_cardpayment', fd, {headers: {'Content-Type': undefined}}).then(function(response) {
	
		   console.log(response);
		}, function errorCallback(response) {
			console.log(response);
	
		});
	}




});

// Angular Factory 傳給後端
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllCardPayment= function(channelID, callback){

    	$http.get('/get_all_cardpayment/'+channelID).success(function(output){
			callback(output)
		});
	}

	factory.queryAllSR= function(callback){

    	$http.get('/get_all_SR/').success(function(output){
			callback(output)
		});
	}

	factory.queryAllStoreData= function(callback){

    	$http.get('/get_all_StoreData/').success(function(output){
			callback(output)
		});
	}

	factory.queryCardPayment = function(id, callback){
    	$http.get('/get_cardpayment/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordCardPayment = function(data,file, callback){


		
		// var cardpayment = data.id + "-" + data.acquiringbank + "-" + data.acquiringdate + "-" + data.size + "-" + data.ipfsHash;
		var fd = new FormData();
		
		fd.append('acquiringbank', data.acquiringbank);
		fd.append('acquiringdate', data.acquiringdate);
		fd.append('file', file);
		fd.append('channel',data.channel );
		console.log(file);
		console.log(data.id);
		console.log(data.acquiringbank);
		// console.log(data.file);
		// console.log(fd['file']);
		$.ajax({

			type: "post",
			url: 'http://localhost:8000/add_cardpayment',
			data: fd,
			cache: false,
			contentType: false,
			processData: false,
		}).done(function (datas) {
			console.log(datas);
			
		})
		// $.ajax({
		// 	method: 'POST',
		// 	url: 'http://localhost:8000/add_cardpayment',
		// 	data: fd,
		// 	headers:{
		// 		'Content-Type': 'multipart/form-data'
		// 	},success:function (respond) {
		// 		console.log("ok");
		// 	}});
	
		// $http.post('/add_cardpayment', fd, {headers: {'Content-Type': 'multipart/form-data'}}).then(function(response) {
		// 	console.log(response);
		//  }, function errorCallback(response) {
		// 	 console.log(response);
	 
		//  });
    	// $http.get('/add_cardpayment/'+cardpayment).success(function(output){
		// 	callback(output)
		// });
	}

	factory.recordSettlementReport = function(data,file, callback){
		
		// var cardpayment = data.id + "-" + data.acquiringbank + "-" + data.acquiringdate + "-" + data.size + "-" + data.ipfsHash;
		var fd = new FormData();
		
		fd.append('reporttype', data.reporttype);
		fd.append('settledate', data.settledate);
		fd.append('settlebank', data.settlebank);
		fd.append('receivebank', data.receivebank);
		fd.append('file',file);
		fd.append('channel',data.channel );
		// console.log(file);
		// console.log(data.id);
		// console.log(data.acquiringbank);
		// console.log(data.file);
		// console.log(fd['file']);
		$.ajax({

			type: "post",
			url: 'http://localhost:8000/add_settlementReport',
			data: fd,
			cache: false,
			contentType: false,
			processData: false,
		}).done(function (datas) {
			console.log(datas);
			
		})
	}

	factory.recordStoreData = function(data,file, callback){
		
		// var cardpayment = data.id + "-" + data.acquiringbank + "-" + data.acquiringdate + "-" + data.size + "-" + data.ipfsHash;
		var fd = new FormData();
		
		fd.append('storeid', data.storeid);
		fd.append('file',file);
		
		$.ajax({

			type: "post",
			url: 'http://localhost:8000/add_StoreData',
			data: fd,
			cache: false,
			contentType: false,
			processData: false,
		}).done(function (datas) {
			console.log(datas);
			
		})
	}

	factory.changeValue = function(data, callback){

		var newvalue = data.id + "-" + data.newsize + "-" + data.newipfshash;

    	$http.get('/change_value/'+newvalue).success(function(output){
			callback(output)
		});
	}

	return factory;
});


