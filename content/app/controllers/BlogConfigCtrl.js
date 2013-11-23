'use strict';


ngBlogApp.controller('BlogConfigCtrl',
    function ($scope, blogConfigService) {
        blogConfigService.fetchConfig(function (d) {
            $scope.Blog = d;
        });
    });