'use strict';


ngBlogApp.controller('PostViewCtrl', function ($location, $scope, $q, $window, $routeParams, postService, authService) {

    $scope.postView = {};

    if ($routeParams.postId) {
        postService.fetchPost($routeParams.postId)
            .then(function (data) {
                $scope.postView.Id = data[0].Id;
                $scope.postView.Title = data[0].Title;
                $scope.postView.Date = data[0].Date;
                $scope.postView.Body = data[0].Body;
                $window.document.title = data[0].Title;
                ;
            }, function (ignored_status) {
                //$console.log(ignored_status);
            });
    } else {
        $location.path('/blog');
    }
});
