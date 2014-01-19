'use strict';


var ngBlogApp = angular.module('ngMicroBlog', ['ngCookies', 'ngRoute', 'chieffancypants.loadingBar'])
    .config(function ($httpProvider, $routeProvider, $locationProvider) {

        $routeProvider.when('/login', {
            templateUrl: 'content/app/directives/login/login.html',
            controller: 'LoginCtrl'
        }).when('/blog', {
                templateUrl: 'content/app/directives/blog/blog.html',
                controller: 'BlogPostsCtrl'
            }).when('/editor', {
                templateUrl: 'content/app/directives/editor/mdedit.html',
                controller: 'MdEditorCtrl'
            }).otherwise({redirectTo: '/blog'});

        $httpProvider.defaults.xsrfHeaderName = 'X-Csrftoken';
        $locationProvider.html5Mode(true);
        //$locationProvider.hashPrefix('!');

    }).run(function ($cookies, $http) {
        $http.defaults.headers.common['X-Csrftoken'] = atob($cookies._xsrf.split('|')[0]);
    });