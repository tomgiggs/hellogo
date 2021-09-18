namespace * manager

enum ExperimentType {
    TIANJI,
    SEARCH,
    RECOMMEND,
    ALGORITHM,
    BMS
}

enum ExperimentPriority {
    HIGH,
    MEDIUM,
    LOW
}

enum ExperimentStatus {
    CREATED,
    RUNNING,
    SUCCESS,
    FAILED,
    DELETED
}

enum TrialStatus{
    CREATED,
    RUNNING,
    SUCCESS,
    FAILED,
    DELETED,
    SUSPENDED,
}

enum TopStandard {
    AUC
}

enum Order{
    ASC,
    DESC
}

enum TianjiType {
    train
    predict
}

enum TrialStatusUpdater {
    controller
    collector
}

//实验
struct ExperimentAddDto {
    //名称
    1:required string Name,
    //描述
    2:string Description,
    //实验类型: 0.天机 1.搜索 2.推荐 3.算法 4.BMS
    3:ExperimentType ExperimentType,
    //实验优先级: 0.top 1.medium 2.low
    4:ExperimentPriority ExperimentPriority,
    //trial并发数
    5:i32 TrialConcurrency,
    //实验最长执行时间
    6:i32 MaxExecDuration,
    //实验中trial的最多数量，达到时实验成功
    7:i32 MaxTrialNum,
    //实验中trial的最多失败数量，达到时实验失败
    8:i32 MaxFailedTrialNum,
    //搜索空间参数
    9:string SearchSpace,
    //优化器配置
    10:string OptimizerConfig,
    //模型路径
    11:string ModelPath,
    //训练容器镜像
    12:string TrialImage,
    //训练容器执行命令
    13:string TrialCmd,
    //初始化容器镜像，用于下载代码和数据(初始化容器相关不用填，目前是硬编码)
    14:string InitContainerImage,
    //初始化容器volume名称
    15:string InitContainerMountName,
    //初始化容器volume路径
    16:string InitContainerMountPath,
    //cpu限制
    17:i32 Cpu,
    //内存限制
    18:i64 Ram,
    //gpu限制
    19:i32 Gpu,
    //gpu内存限制
    20:double GpuRam,
    //数据路径
    21:string DataSetPath,
    //代码路径
    22:string CodePath,
    //最优trial个数
    23:i32 TopNum,
    //最优trial的标准，0.AUC
    24:TopStandard TopStandard,
    //topk取最小还是最大，0.最小 1.最大
    25:Order TopOrder,
    //天机项目ID
    26:i64 TianjiProjectId,
    //天机用户ID
    27:i64 TianjiUserId,
    //天机模型版本
    28:i64 TianjiModelVersion,
    //天机实验类型，0.train 1.predict
    29:TianjiType TianjiType,
    //天机产出模型名
    30:string TianjiUsedModel,
    //天机存放在ceph的bucket
    31:string TianjiDatasetBucket,
    //天机离线预测的id
    32:i64 TianjiPredictID,
    //虚假的实验ID
    33:i64 FkExperimentId,
    //单个trial最大执行时间，单位秒
    34:i64 MaxTrialExecDuration,
    //不对接metaai
    35:bool NoSuggestion
    //对数据集的注释
    36:string AnnotationData
    //使用集成模型
    37:bool UseEnsemble
}

//实验更新
struct ExperimentEditDto{
    1:required i64 Id,
    2:string StartTime
    3:string EndTime,
    4:i64 ExecDuration,
    5:ExperimentStatus Status
}

//trial上报指标
struct TrialEditDto {
    1:required i64 Id,
    2:required string TrialName,
    3:string ModelName,
    4:string UsedModels,
    5:string RecomModel,
    6:string MetricType,
    7:double MetricValue,
    8:string ModelPath,
    9:string ModelParams,
    10:string FeatureNum,
    11:string FeatureImportance,
    12:string EvalResult,
    13:string PredictProb,
    14:string EvalData,
    15:string StartTime,
    16:string EndTime,
    17:i64 Duration,
    18:TrialStatus Status,
    19:i64 ModelVersion,
    20:string TrainConsuming,
    21:string ParamSpace,
    22:double BestMetric,
    23:string NodeIP,
    24:string PerformanceType,
    25:double PerformanceValue,
    26:string PreviewData,
    27:string PredictResultFileName,
    28:TrialStatusUpdater Updater
}

//推理
struct InferAddDto {
    //用户id
    1:required i64 UserID,
    //项目id
    2:required i64 ProjectID,
    //名称
    3:required string Name,
    //服务ID
    4:required i64 ServiceID,
    //服务镜像
    5:required string ServiceImage,
    //服务端口
    6:required i32 ServicePort,
    //服务命令
    7:required string ServiceCmd,
    //服务pod数量
    8:required i32 ServicePodNum,
    //服务url的路由
    9:required string ServiceRouterPath,
    //cpu限制
    10:i32 Cpu,
    //内存限制
    11:i64 Ram,
    //gpu限制
    12:i32 Gpu,
    //gpu内存限制
    13:double GpuRam,
    //使用模型
    14:required string UsedModel,
    //模型路径
    15:required string ModelPath,
    //天机存放在ceph的bucket
    16:required string Bucket
}

//异常格式
exception ManagerException {
    1:i32 errorCode,
    2:string message,
}

service DbManager {
    //创建实验
    i64 CreateExperiment(1:ExperimentAddDto experiment) throws (1:ManagerException err),
    //更新实验状态
    bool UpdateExperiment(1:ExperimentEditDto experiment) throws (1:ManagerException err),
    //删除实验
    bool DeleteExperiment(1:i64 experimentId) throws (1:ManagerException err),
    //创建trial记录
    i64 CreateTrial(1:i64 experimentId,2:i32 trialNo,3:string trialName) throws (1:ManagerException err),
    //暂停trial
    bool SuspendTrial(1:i64 id) throws (1:ManagerException err),
    //恢复trial
    bool ResumeTrial(1:i64 id) throws (1:ManagerException err),
    //删除trial
    bool DeleteTrial(1:i64 id) throws (1:ManagerException err),
    //更新trial状态
    bool UpdateTrialStatus(1:TrialEditDto trial) throws (1:ManagerException err),
    //创建推理服务
    i64 CreateInfer(1:InferAddDto infer) throws (1:ManagerException err),
    //暂停推理服务
    bool SuspendInfer(1:i64 id) throws (1:ManagerException err),
    //恢复推理服务
    bool ResumeInfer(1:i64 id) throws (1:ManagerException err),
    //删除推理服务
    bool DeleteInfer(1:i64 id) throws (1:ManagerException err),
}