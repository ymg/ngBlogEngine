'use strict';


ngBlogApp.controller('DbCfgCtrl', function ($scope, $window, $location, $log, blogConfigService) {

    $window.document.title = "Blog Configuration";

    $scope.dbCfg = {};

    $scope.updateConfig = function () {
        if ($scope.dbCfg.$valid) {

        }
    };

    blogConfigService.fetchDbConfig()
        .then(function (d) {
            $scope.dbCfg = d;
        }, function (status) {
            $log.error(status);
        });


});
