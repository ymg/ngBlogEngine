'use strict';


ngBlogApp.controller('PostViewCtrl', function ($location, $scope, $q, $window, $routeParams, postService, authService) {

    $scope.postView = {};

    if ($routeParams.postId) {
        postService.fetchPost($routeParams.postId)
            .then(function (data) {

                $scope.postView.Id = data.Id;
                $scope.postView.Title = data.Title;
                $scope.postView.Date = data.Date;
                $scope.postView.Body = data.Body;
                $window.document.title = data.Title;

            }, function (ignored_status) {
                //$console.log(ignored_status);
            });
    } else {
        $location.path('/blog');
    }
});
