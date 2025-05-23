<?php

namespace ccxt\async;

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code

use Exception; // a common import
use ccxt\async\abstract\huobi as htx;

class huobi extends htx {

    public function describe(): mixed {
        return $this->deep_extend(parent::describe(), array(
            'id' => 'huobi',
            'alias' => true,
        ));
    }
}
