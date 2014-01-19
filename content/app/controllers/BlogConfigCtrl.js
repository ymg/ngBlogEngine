'use strict';


ngBlogApp.controller('BlogConfigCtrl',
    function ($scope, blogConfigService) {
        blogConfigService.fetchConfig()
            .then(function (data) {
                console.log(data);
                $scope.blog = data;
            }, function (status) {
                console.log(status);
            });
    });