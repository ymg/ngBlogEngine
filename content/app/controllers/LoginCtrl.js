'use strict';


ngBlogApp.controller('LoginCtrl', function ($rootScope, $scope, $location, $window, authService) {
    $scope.cred = {}
    $scope.login = function () {
        if ($scope.loginform.$valid) {
            authService.auth(this.cred, function (stat) {
                if (stat === 200) {
                    $rootScope.loginStatus = authService.isLoggedIn();
                    $location.path('/blog');
                } else {
                    $.UIkit.notify({
                        message: 'Unable to authenticate your credentials',
                        timeout: 3000,
                        pos: 'top-center'
                    });
                    $location.path('/login');
                }
            });
            $scope.cred = {}
            $scope.loginform.$setPristine();
        }
    };
    $window.document.title = 'Admin Login';
});