'use strict';


ngBlogApp.controller('DbCfgCtrl', function ($scope, $window, $location) {

    $window.document.title = "Blog Configuration";

    $scope.blogCfg = {};

    $scope.updateConfig = function () {
        if ($scope.blogCfg.$valid) {

        }
    };


});
