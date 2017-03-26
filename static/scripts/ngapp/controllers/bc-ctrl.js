(function() {

    var appName = "ball-clock";
    var ctrlName = "bc-ctrl";

    var app = angular.module(appName);

    app.controller(ctrlName, function($scope, $http) {

        $scope.runSimulation = function() {
            $scope.showContent("running");
            //$scope.clock.interval = parseInt($scope.clock.interval);
            $http.post(window.location.origin + "/RunSim/", JSON.stringify($scope.clock))
                .then($scope.results, null);

        };

        $scope.clock = {
            interval: 5,
            ballCount: 30,
            years: 0,
            months: 0,
            days: 0,
            hours: 0,
            minutes: 0
        };

        $scope.minCount = 0;
        $scope.maxCount = 0;

        $scope.updateMinMaxCount = function() {
            switch($scope.clock.interval){
                case "5":
                    $scope.minCount = 29;
                    $scope.maxCount = 127;
                    break;
                case "10":
                    $scope.minCount = 27;
                    $scope.maxCount = 127;
                    break;
                default:
                    $scope.minCount = 29;
                    $scope.maxCount = 127;
                    break;
            }
            $scope.$apply();
        }

        $scope.getNumber = function(num) {
            var array = new Array();
            for (var i = 0 ; i < num; i++){
                array.push(i);
            }
            return array;
        };

        $scope.start = function() {
            $scope.showContent("sim1");

        };

        $scope.totalMinutesCounted = function() {
            var total = 0;
            for(var i = 0, len =  $scope.result.data.Balls.Balls.length; i < len; i++){
                total = total + $scope.result.data.Balls.Balls[i].PickupCount;
            }
            return total;
        };



        $scope.levelArrays = {
            l1: [],
            l2: [],
            l3: [],
            tc: 0
        };

        $scope.results = function(data, status, headers, config) {

            $scope.result = data;
            $scope.levelArrays.l1 = $scope.getNumber($scope.result.data.Levels[0].MaxBalls);
            $scope.levelArrays.l2 = $scope.getNumber($scope.result.data.Levels[1].MaxBalls);
            $scope.levelArrays.l3 = $scope.getNumber($scope.result.data.Levels[2].MaxBalls);
            var ct = $scope.result.data.RoundTime;
            if (ct > 0){
                $scope.cycleTime = ct;
            }else{
                $scope.cycleTime = "Unreached During Simulation";
            };
            $scope.levelArrays.tc = $scope.totalMinutesCounted();
            $scope.showContent("results");
            $scope.$apply();


        };

        $scope.about = function() {
            $scope.showContent("home");
        };

        $scope.result = {};

        $scope.showContent = function(content){
             switch(content) {
                case "running":
                    $("#running").show();
                    $("#home").hide();
                    $("#sim1").hide();
                    $("#results").hide();
                    break;
                case "home":
                    $("#running").hide();
                    $("#home").show();
                    $("#sim1").hide();
                    $("#results").hide();
                    break;
                case "sim1":
                    $("#running").hide();
                    $("#home").hide();
                    $("#sim1").show();
                    $("#results").hide();
                    break;
                case "results":
                    $("#running").hide();
                    $("#home").hide();
                    $("#sim1").hide();
                    $("#results").show();
                    break;
            }
        }


        $scope.about();


    });

}());