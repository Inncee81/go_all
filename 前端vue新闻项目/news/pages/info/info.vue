<template>
	<view class="ls">
		<view class="t2">{{title}}</view>
			<view class="gd">
		<view class="">{{author_name}}</view>
		<view class="left">{{created_at}}</view>
		</view>
		<rich-text class="richText text" :nodes="strings"></rich-text>
	</view>
</template>

<script>
	export default {
onLoad(e) {
	console.log(e.newid);
	uni.request({
		url: 'https://unidemo.dcloud.net.cn/api/news/36kr/'+e.newid,
		method: 'GET',
		data: {},
		success: res => {
			console.log(res.data);
			this.title=res.data.title
			this.strings=res.data.content
			this.created_at=res.data.created_at
			this.author_name=res.data.author_name
		},
		fail: () => {},
		complete: () => {}
	});
}
,		data() {
			return {
				title:'',
				strings:'',
				created_at:'',
				author_name:''
			}
		},
		methods: {
			
		}
	}
</script>

<style>
 .t2{
					font-weight:bold;
					 line-height: 70upx;
				 	 font-size: 54upx;
										 text-align: justify;
				 }
				 .text{
					 text-align: justify;
					 FONT-SIZE:36upx;
					 LINE-HEIGHT:54upx;
				 }
.ls{
					
					margin-left: 20upx;
					margin-right: 20upx;
				}
				
				.gd{
									 display: flex;
									 flex-direction: row;
									 	justify-content:flex-start;
									color: #777777;
									margin-top: 20upx;
									margin-bottom: 20upx;
				}
				.left{
									 padding-left:30upx ;
				}
</style>
