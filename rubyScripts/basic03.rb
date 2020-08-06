HeaderConfig = Hash.new('HeaderConfig')
HeaderConfig['Login']=['mod_name','log_date','log_time','channelcode','app_version','ip','platform','equdid','imei','mac','mobilemodel','screen_x','screen_y','mobileosversion','netmode','sessionid','iccid','imsi','serverid','accountname','accountid','serveraccountid','userid','channelaccountid','sub_channelcode','level','battle_lev','vip_lev'],
HeaderConfig['Logout']=['mod_name','log_date','log_time','channelcode','platform','equdid','session_id','imei','mac','serverid','accountname','accountid','sub_channelcode','level'],
HeaderConfig['Card']=['mod_name','log_date','log_time','channelcode','accountid','userid','','platform','number','itemcount','systemtype','type','voucher','sub_channelcode'],
HeaderConfig['EmoneyCost']=['mod_name','log_date','log_time','type','userid','number','id_target','itemtype','itemcount','price','reservel','reserve2','reserve3'],
HeaderConfig['Item']=['mod_name','log_date','log_time','type','userid','itemtype','itemnum'],
HeaderConfig['Task']=['mod_name','log_date','log_time','type','userid','taskid'],
HeaderConfig['Box']=['mod_name','log_date','log_time','userid','itemtype','result'],
HeaderConfig['Level']=['mod_name','log_date','log_time','type','userid','levelid','time'],
HeaderConfig['PVE']=['mod_name','log_date','log_time','battleid','userid','type','result','array','monstergroup','time'],
HeaderConfig['PVP']=['mod_name','log_date','log_time','battleid','leftuserid','rightuserid','type','result','leftarray','rightarray','time'],
HeaderConfig['Activity']=['type','data1','data2'],
HeaderConfig['Newplayer']=['mod_name','log_date','log_time','type','userid'],
HeaderConfig['First_level']=['mod_name','log_date','log_time','type','userid','data'],
HeaderConfig['Special_battle']=['mod_name','log_date','log_time','battleid','userid','result','array','monstergroup','time'],
HeaderConfig['Loading']=['mod_name','log_date','log_time','type','userid'],
HeaderConfig['Strengthen']=['mod_name','log_date','log_time','type','userid','data1','data2'],
HeaderConfig['Pay']=['mod_name','log_date','log_time','userid'],
HeaderConfig['Online']=['mod_name','log_date','log_time','serverid','platform','number'],
HeaderConfig['Shop']=['mod_name','log_date','log_time','userid','type','itemtype','num','money_type','money_num'],
HeaderConfig['syn']=['mod_name','log_date','log_time','type','userid','data1','data2'],
HeaderConfig['Poker']=['mod_name','log_date','log_time','type','data1','data2','data3'],
HeaderConfig['Coin']=['mod_name','log_date','log_time','userid','type','num'],
HeaderConfig['luckymoney']=['mod_name','log_date','log_time','type','data1','data2','data3']
  def packData (header,data)
    if header.length == data.length then
      for i in 0...header.length
        puts header[i],data[i]
        # event.set(header[i],data[i])
      end
    end
  
  end  
#  puts obj['Coin']
# puts obj.keys
modName = 'Activity'
logInfo = '13,18257[1559][10],10126[0][0]'
for key in HeaderConfig.keys do
  # puts "key is:"+key
  if key == 'Activity' then
    # puts obj[key]
    packData(HeaderConfig[key],logInfo.split(','))
  end
end
