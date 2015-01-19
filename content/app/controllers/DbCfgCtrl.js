'use strict';


ngBlogApp.controller('DbCfgCtrl', function ($scope, $window, $location, $log, blogConfigService) {

    $window.document.title = "Blog Configuration";

    $scope.cfg = {};

    $scope.updateConfig = function () {
        if ($scope.dbcfg.$valid) {
            blogConfigService.updateDbConfig(this.dbcfg).then(function (status) {

            }, function (status) {

            });

            $scope.cfg = {};
        }
    };

    blogConfigService.fetchDbConfig()
        .then(function (d) {
            $scope.cfg = d;
        }, function (status) {
            $log.error(status);
        });


});
