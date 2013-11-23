'use strict';


ngBlogApp.service('postService', function ($log, $http) {
    this.fetchPosts = function () {
        return {"Title": "ngMicroBlog!", "Body": "Hello, ngMicroBlog", "Date": "17-Nov-2013"};
        /*$http({method: 'GET', url: 'http://127.0.0.1:8080/'}).
         success(function(data, status, headers, config) {
         return succ(data);
         }).
         error(function(data, status, headers, config) {
         $log.warn(data, status, headers, config);
         });*/
    };
});


