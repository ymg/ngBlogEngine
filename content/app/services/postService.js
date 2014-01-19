'use strict';


ngBlogApp.service('postService', function ($log, $http) {
    this.fetchPosts = function () {
        return {"Title": "ngMicroBlog!", "Body": "Hello, ngMicroBlog", "Date": "17-Nov-2013"};
    };
});


