'use strict';


ngBlogApp.controller('LoginCtrl',
    function ($scope, postService, $window, authenticateAdminService) {
        authenticateAdminService.checkSession(function(status) {
            if(status === 200){
                $scope.postList = postService.fetchPosts();
                $window.document.title = 'Admin Login';
            }
        });
    });