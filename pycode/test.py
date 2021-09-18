#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Desc   :
import numpy as np

from GreedySelect import GreedySelect


def prepare_data():
    trials_prob = [
        [  # 预测概率
            0.3174,
            0.3201,
            0.887,
            0.3127
        ],
        [  # 预测概率
            0.3174,
            0.3201,
            0.887,
            0.3127
        ],
        [  # 预测概率
            0.3174,
            0.3201,
            0.887,
            0.3127
        ],
        [  # 预测概率
            0.3174,
            0.3201,
            0.887,
            0.3127
        ]
    ]

    trial_label = [  # 实际正负值，0或1
        0,
        0,
        1,
        1
    ]
    return trials_prob, trial_label


def main(xx):
    trials_prob, trial_label = prepare_data()

    # 开始
    trials_prob, trial_label = np.array(trials_prob), np.array(trial_label)
    gs = GreedySelect(max_trial_num=2, metric="auc")
    indices = gs.fit(val_preds=trials_prob, val_label=trial_label)

    idx_trials_prob = trials_prob[indices]
    ensem_prob = gs.predict(idx_trials_prob)
    print(ensem_prob)

    metric = gs.calcu_metric(trial_label, ensem_prob, metric=gs.metric)
    print(f"metric {gs.metric}:{metric}")
    return '{}:{}'.format({gs.metric}, {metric})


if __name__ == "__main__":
    main()
