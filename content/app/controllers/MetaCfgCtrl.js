'use strict';


ngBlogApp.controller('MetaCfgCtrl', function ($location, $scope, $q, $log, $window, $routeParams, authService) {

    $scope.metacfgform = {};

    $scope.saveChanges = function () {
        if ($scope.metaform.$valid) {
            authService.updateCredentials(this.metacfgform).then(function (stat) {
                if (stat === 200) {
                    $.UIkit.notify({
                        message: 'Password Updated',
                        timeout: 3000,
                        pos: 'top-center',
                        status: 'success',
                    });
                    $location.path('/blog');
                } else {
                    $.UIkit.notify({
                        message: 'Unable to process your request at this time',
                        timeout: 3000,
                        pos: 'top-center',
                        status: 'danger',
                    });
                    $location.path('/login');
                }
            }, function (stat) {
                $.UIkit.notify({
                    message: 'Unauthorized Request',
                    timeout: 3000,
                    pos: 'top-center',
                });
            });
        }
    };
    $window.document.title = 'Update Admin Password';
});
