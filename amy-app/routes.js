//SPDX-License-Identifier: Apache-2.0

var cardpayment = require('./controller.js');

module.exports = function(app){

  //app.METHOD(PATH, HANDLER) METHOD 是 HTTP 要求方法。PATH 是伺服器上的路徑。HANDLER 是當路由相符時要執行的函數
  app.get('/get_cardpayment/:id', function(req, res){
    cardpayment.get_cardpayment(req, res);
  });
  app.post('/add_cardpayment', function(req, res){
    cardpayment.add_cardpayment(req, res);
  });
  app.get('/get_all_cardpayment/:channelID', function(req, res){
    cardpayment.get_all_cardpayment(req, res);
  });
  app.get('/change_value/:newvalue', function(req, res){
    cardpayment.change_value(req, res);
  });
  app.get('/delete_cardpayment/:id', function(req, res){
    cardpayment.delete_cardpayment(req, res);
  });
  app.get('/ipfs_download/:ipfshash', function(req, res){
    cardpayment.ipfs_download(req, res);
  });
  app.post('/add_settlementReport', function(req, res){
    cardpayment.add_settlementReport(req, res);
  });
  app.get('/get_all_SR', function(req, res){
    cardpayment.get_all_SR(req, res);
  });
  app.post('/add_StoreData', function(req, res){
    cardpayment.add_StoreData(req, res);
  });
  app.get('/get_all_StoreData', function(req, res){
    cardpayment.get_all_StoreData(req, res);
  });
  app.post('/upload_file', function(req, res){
    cardpayment.upload_file(req, res);
  });
}
