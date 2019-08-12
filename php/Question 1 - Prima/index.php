<?php

function Solution($len) {
    $result = array();
    for ($i=1; $i <= $len; $i++) { 
        $a = 0;
        for ($j=1; $j <= $i ; $j++) { 
            if ($i % $j == 0) {
                $a++;
            }
        }
        if ($a == 2) {
            $result[] = $i;
        }
    }
    return $result;
}

print_r(Solution(15));
