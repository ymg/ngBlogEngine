'use strict';


ngBlogApp.controller('AuthCtrl',
    function (authenticateAdminService, $scope, $location) {
        $scope.cred = {}
        $scope.login = function () {
            if ($scope.loginform.$valid) {
                authenticateAdminService.auth(this.cred, function (stat) {
                    if (stat === 200) {
                        $.UIkit.notify({
                            message : 'Bazinga!',
                            status  : 'info',
                            timeout : 3000,
                            pos     : 'top-center'
                        });
                        $location.path('/blog');
                    }
                });

                $scope.cred = {}
                $scope.loginform.$setPristine();
            }
        }
    });