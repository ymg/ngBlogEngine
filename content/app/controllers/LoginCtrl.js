'use strict';


ngBlogApp.controller('LoginCtrl',
    function (postService, $scope) {
        $scope.postList = postService.fetchPosts();
    });