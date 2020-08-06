require 'json'
HeaderConfig = JSON.parse('{
    "Login": ["channelcode","app_version","ip","platform","equdid","imei","mac","mobilemodel","screen_x","screen_y","mobileosversion","netmode","sessionid","iccid","imsi","serverid","accountname","accountid","serveraccountid","userid","channelaccountid","sub_channelcode","level","battle_lev","vip_lev"],
    "Logout": ["channelcode","platform","equdid","session_id","imei","mac","serverid","accountname","accountid","sub_channelcode","level"],
    "Card": ["channelcode","accountid","userid","","platform","number","itemcount","systemtype","type","voucher","sub_channelcode"],
    "EmoneyCost": ["type","userid","number","id_target","itemtype","itemcount","price","reservel","reserve2","reserve3"],
    "Item": ["type","userid","itemtype","itemnum"],
    "Task": ["type","userid","taskid"],
    "Box": ["userid","itemtype","result"],
    "Level": ["type","userid","levelid","time"],
    "PVE": ["battleid","userid","type","result","array","monstergroup","time"],
    "PVP": ["battleid","leftuserid","rightuserid","type","result","leftarray","rightarray","time"],
    "Activity": ["type","data1","data2"],
    "Newplayer": ["type","userid"],
    "First_level": ["type","userid","data"],
    "Special_battle": ["battleid","userid","result","array","monstergroup","time"],
    "Loading": ["type","userid"],
    "Strengthen": ["type","userid","data1","data2"],
    "Pay": ["userid"],
    "Online": ["serverid","platform","number"],
    "Shop": ["userid","type","itemtype","num","money_type","money_num"],
    "syn": ["type","userid","data1","data2"],
    "Poker": ["type","data1","data2","data3"],
    "Coin": ["userid","type","num"],
    "luckymoney": ["type","data1","data2","data3"]
  }')

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

