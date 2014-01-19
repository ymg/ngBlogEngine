'use strict';


ngBlogApp.controller('BlogPostsCtrl',
    function ($scope, $window) {
        hlight();
        $window.document.title = 'Blog';
    });