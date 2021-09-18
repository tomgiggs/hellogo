#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Desc   :

import numpy as np
from sklearn import metrics
from EnsembleBase import EnsembleBase


class GreedySelect(EnsembleBase):
    def fit(self, val_preds: np.ndarray, val_label: np.ndarray):
        assert self.max_trial_num <= len(val_preds)
        trajectory = []  # 用于记录每次添加 trial 后的融合结果对应得分
        ensemble = []  # 用于记录被选中 trials 的预测结果
        order = []  # 用于记录被选中 trials 的 trial_id

        for i in range(self.max_trial_num):  # 遍历往 ensemble 中添加 max_trial_num 个 trial 预测结果

            # 计算已经在 ensemble 中被选中的 trial 的预测结果平均值
            scores = np.zeros(len(val_preds))
            s = len(ensemble)
            if s == 0:
                weighted_ensemble_prediction = np.zeros(val_preds[0].shape)
            else:
                ensemble_prediction = np.zeros(ensemble[0].shape)
                for pred in ensemble:
                    ensemble_prediction += pred

                ensemble_prediction /= s
                weighted_ensemble_prediction = (s / float(s + 1)) * ensemble_prediction

            # 遍历所有 trial 的预测结果，并计算每个 trial 与当前 ensemble 被选中的 trials 的融合结果平均值和对应指标得分
            fant_ensemble_predicion = np.zeros(weighted_ensemble_prediction.shape)
            for j, pred in enumerate(val_preds):
                fant_ensemble_predicion[:] = weighted_ensemble_prediction + (1. / float(s + 1)) * pred
                scores[j] = self.calcu_metric(y_true=val_label, y_pred=fant_ensemble_predicion, metric=self.metric)

            # 对scores 排序，找到对于当前 ensemble 融合效果最好的 trial， 记录 id 和得分
            scores_sort = np.argsort(scores)
            best = [sid for sid in scores_sort if sid not in order][0]
            ensemble.append(val_preds[best])  # 往 ensemble 中添加预测结果
            trajectory.append(scores[best])  # 往 trajectory 中添加当前 ensemble 的得分
            order.append(best)  # 往 order 中添加 trial_id

        # 找到 trajectory 中最优得分对应的 index，并从 order 找到所需要的 trial_id， 返回结果
        min_score = np.min(trajectory)
        first_index_of_best = trajectory.index(min_score)
        self.indices = indices = order[: first_index_of_best + 1]
        return indices

    def predict(self, test_preds):
        ensemble_pred = None
        for iid, y_pred in enumerate(test_preds):

            if iid == 0:
                ensemble_pred = y_pred / len(test_preds)
            else:
                ensemble_pred += y_pred / len(test_preds)
        return ensemble_pred

    def calcu_metric(self, y_true, y_pred, metric):
        if metric == 'f1':
            score = 1 - metrics.f1_score(y_true, np.argmax(y_pred, axis=1))
            return score

        if metric == 'auc':
            score = 1 - metrics.roc_auc_score(y_true, y_pred)
            return score

        if metric == 'mse':
            score = metrics.mean_squared_error(y_true, y_pred)
            return score
