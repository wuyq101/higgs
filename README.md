# higgs

## 使用外部抽取器来抽取网页具体内容

### 抓取百度搜索内容的例子
1. 在config文件中增加 "ExtractorService" "ExtractorService": "http://localhost:5000/extract/",

2. 在config的Templates下增加 test:“test.json”

3. 创建 test.json文件, external_extractor是外部的抽取器，extractor是higgs自带的go抽取器，如果配置了外部抽取器，优先使用外部的


	`{
  		"disable_out_pub_key": true,
  		"steps": 
  		[{
      		"need_param": "query",
      		"page": "http://www.baidu.com/s?wd={{.query}}",
      		"method": "GET",
      		"header":
      		 {
        		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) 	AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.130 Safari/537.36",
        		"Referer": "https://www.baidu.com/",
        		"Host": "www.baidu.com",
        		"Cookie": "BAIDUID=8C9B6A45045AFE64B00BCFB961A33495:SL=0:NR=50:FG=1"
      		},
      		"output_filename": "index.html",
      		"external_extractor":
      		 {
        		"method": "baidu"
      		},
      		"extractor":
      		 {
        		"search":
        		 {
          			"_root": "#content_left .c-container",
          			"_array": true,
          			"summary": "div.c-abstract",
          			"title": "h3.t a",
          			"result_link": "div.f13 a.c-showurl"
        		}
     		}
   		}]
	}`
	
	在这个配置文件中，外部抽取方法名为 baidu， 所以爬虫会将网页内容post到如下地址 http://localhost:5000/extract/baidu ,这要这个url能正确返回抽取内容就可以, 这里假如返回的是测试内容具体是一个json: {\"test\":\"这里是测试抽取的内容\"}
	
	
4. 提交爬虫请求: 根据配置文件， 抓百度需要一个参数query，加入我们提交的词是“李世明”， 发送GET 请求到 http://localhost:8001/submit?query=李世明&tmpl=test
5. 然后爬虫就会根据配置去抓取 http://www.baidu.com/s?wd=李世明 的网页内容，并且将网页内容存到/Users/wuyq/higgs/data/test/2016/08/08/1470643274095181972/index.html， 同时会将网页内容post到外部抽取器，这里假设外部抽取器返回测试内容 {\"test\":\"这里是测试抽取的内容\"}
6. http://localhost:8001/submit?query=李世明&tmpl=test 这个请求最后的返回内容如下：   
	`
	{
    	"status": "finish_fetch_data",
    	"need_param": "",
    	"id": "test|2016|08|08|1470643274095181972",
    	"data": "{\"test\":\"这里是测试抽取的内容\"}"
	}
	`
7. 如果不配置external_extractor， 那么根据配置文件中extractor的规则配置，抽取搜索页的具体内容
	
	
