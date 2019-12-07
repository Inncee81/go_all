<template>
<view class="list">
	<view class="ls"  v-for="(item,index) in news" :key="index" @tap="openinfo" :data-newid="item.post_id">
	<view class="wl ">
		<view class="t2">{{item.title| qstr}}</view>
		<view class="gd">
			
		<view class="t1">{{item.author_name}}</view>
		<view class="t1 left">{{item.published_at|qtime}}</view>
		</view>
	</view>		
	
	<view class="wh ">
		
		<image class="img" :src="item.author_avatar"></image>
	</view>	
</view>		
</view>		
</template>

		
<script>
	export default {
		data() {
			return {
			news:[],
			code:''
			}
		},
		onLoad() {
uni.request({
	url: 'https://unidemo.dcloud.net.cn/api/news',
	method: 'GET',
	data: {},
	success: res => {
		//console.log(res.data);
		this.news=res.data;
		this.code=res.data;
		
	},
	fail: () => {},
	complete: () => {}
});
		},
		methods: {
	openinfo(e){
		//console.log(e);
var newid=e.currentTarget.dataset.newid
//console.log(newid);
		uni.navigateTo({
			url: '../info/info?newid='+newid,
			success: res => {},
			fail: () => {},
			complete: () => {}
		});
}
		},
		 filters: {
			 qtime: function (value) {
				 
				 return  value.split(" ")[0]
			 } ,
			 qstr: function (str) {
			 				 
							 if(str.length > 30){
								str = str.substring(0,30) +"..."; 
							 } 
			 				 return  str
			 } 
		 }
	}
</script>

<style>
	.list{
			/* border:1px solid #000000; */
				}	
				.ls{
					display: flex;
					flex-direction: row;
						justify-content:space-between;
					margin-left: 20upx;
					margin-right: 20upx;
				}
				 .wl{
					 display: flex;
					 flex-direction: column;
					 	justify-content:space-between;
					 width: 460upx;
					 height: 200upx;
					  /* border:1px solid #000000; */
					  margin-top: 20upx;
					 margin-bottom: 20upx;
				 }
				 .wh{
				 				width: 220upx;
				 				 height: 200upx;
								 
				 				  /* border:1px solid #000000; */
				 				  margin-top: 20upx;
				 				 margin-bottom: 20upx;
				 }
					 
				 .gd{
					 display: flex;
					 flex-direction: row;
					 	justify-content:flex-start;
				 }
					 
				 .img{
					     max-width:100%;  
						 max-height: 100%;
				 }
				 .t1{
					
					 font-size: 24upx;
				 }
				 .t2{
					font-weight:bold;
					 line-height: 52upx;
				 	 font-size: 34upx;
					 text-align: justify;
				 }
				 .left{
					 padding-left:30upx ;
				 }
				 .text{
				 					 text-align: justify;
				 }
</style>
