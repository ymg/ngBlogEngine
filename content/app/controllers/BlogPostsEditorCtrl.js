'use strict';


ngBlogApp.controller('BlogPostsEditorCtrl', function ($scope, $window, $location, postService) {
    $window.document.title = 'Post Editor';

    $scope.InitMarkDownEditor();

    $scope.postForm = {};
    $scope.postNew = function () {
        if ($scope.postform.$valid) {
            var markdownEditor = angular.element('.CodeMirror')[0].CodeMirror;
            this.postForm.markdown = markdownEditor.getValue();
            postService.createNew(this.postForm)
                .then(function (status) {
                    if (status === 200) {
                        $scope.postForm = {};
                        markdownEditor.setValue('');
                        $scope.postform.$setPristine();
                    } else {
                        $location.path('/login');
                    }
                },
                function (ignored_status) {
                    $.UIkit.notify({
                        message: 'Failed processing your request',
                        timeout: 3000,
                        pos: 'top-center'
                    });
                });
        }
    };

});