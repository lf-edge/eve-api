# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: info/edge_node_cluster.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1cinfo/edge_node_cluster.proto\x12\x13org.lfedge.eve.info\x1a\x1fgoogle/protobuf/timestamp.proto\"\x94\x01\n\x11KubeNodeCondition\x12\x38\n\x04type\x18\x01 \x01(\x0e\x32*.org.lfedge.eve.info.KubeNodeConditionType\x12\x0b\n\x03set\x18\x02 \x01(\x08\x12\x38\n\x14last_transition_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xba\x02\n\x0cKubeNodeInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12:\n\nconditions\x18\x02 \x03(\x0b\x32&.org.lfedge.eve.info.KubeNodeCondition\x12\x13\n\x0brole_server\x18\x03 \x01(\x08\x12\x36\n\x12\x63reation_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1a\n\x12\x61pi_server_sersion\x18\x05 \x01(\t\x12\x13\n\x0binternal_ip\x18\x06 \x01(\t\x12\x13\n\x0bschedulable\x18\x07 \x01(\x08\x12<\n\x10\x61\x64mission_status\x18\x08 \x01(\x0e\x32\".org.lfedge.eve.info.NodeAdmission\x12\x0f\n\x07node_id\x18\t \x01(\t\"\xa4\x01\n\x14KubePodNameSpaceInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tpod_count\x18\x02 \x01(\r\x12\x19\n\x11pod_running_count\x18\x03 \x01(\r\x12\x19\n\x11pod_pending_count\x18\x04 \x01(\r\x12\x18\n\x10pod_failed_count\x18\x05 \x01(\r\x12\x1b\n\x13pod_succeeded_count\x18\x06 \x01(\r\"\x82\x02\n\x11KubeEVEAppPodInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x32\n\x06status\x18\x02 \x01(\x0e\x32\".org.lfedge.eve.info.KubePodStatus\x12\x15\n\rrestart_count\x18\x03 \x01(\r\x12\x35\n\x11restart_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x36\n\x12\x63reation_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x12\n\nip_address\x18\x06 \x01(\t\x12\x11\n\tnode_name\x18\x07 \x01(\t\"\x9f\x01\n\x15KubeVolumeReplicaInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x12\n\nowner_node\x18\x02 \x01(\t\x12#\n\x1brebuild_progress_percentage\x18\x03 \x01(\r\x12?\n\x06status\x18\x04 \x01(\x0e\x32/.org.lfedge.eve.info.StorageVolumeReplicaStatus\"\xdc\x03\n\x0eKubeVolumeInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x36\n\x05state\x18\x02 \x01(\x0e\x32\'.org.lfedge.eve.info.StorageVolumeState\x12@\n\nrobustness\x18\x03 \x01(\x0e\x32,.org.lfedge.eve.info.StorageVolumeRobustness\x12\x36\n\x12\x63reation_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x19\n\x11provisioned_bytes\x18\x05 \x01(\x04\x12\x17\n\x0f\x61llocated_bytes\x18\x06 \x01(\x04\x12?\n\npvc_status\x18\x07 \x01(\x0e\x32+.org.lfedge.eve.info.StorageVolumePVCStatus\x12;\n\x07replica\x18\x08 \x03(\x0b\x32*.org.lfedge.eve.info.KubeVolumeReplicaInfo\x12\x45\n\x13robustness_substate\x18\t \x01(\x0e\x32(.org.lfedge.eve.info.StorageHealthStatus\x12\x11\n\tvolume_id\x18\n \x01(\t\"\xb0\x01\n\x0fKubeStorageInfo\x12\x32\n\x06health\x18\x01 \x01(\x0e\x32\".org.lfedge.eve.info.ServiceStatus\x12\x33\n\x0ftransition_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x34\n\x07volumes\x18\x03 \x03(\x0b\x32#.org.lfedge.eve.info.KubeVolumeInfo\"\xe1\x01\n\x0bKubeVMIInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x32\n\x06status\x18\x02 \x01(\x0e\x32\".org.lfedge.eve.info.KubeVMIStatus\x12\x31\n\rcreation_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x38\n\x14last_transition_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x10\n\x08is_ready\x18\x05 \x01(\x08\x12\x11\n\tnode_name\x18\x06 \x01(\t*\x98\x02\n\x15KubeNodeConditionType\x12(\n$KUBE_NODE_CONDITION_TYPE_UNSPECIFIED\x10\x00\x12\"\n\x1eKUBE_NODE_CONDITION_TYPE_READY\x10\x01\x12,\n(KUBE_NODE_CONDITION_TYPE_MEMORY_PRESSURE\x10\x02\x12*\n&KUBE_NODE_CONDITION_TYPE_DISK_PRESSURE\x10\x03\x12)\n%KUBE_NODE_CONDITION_TYPE_PID_PRESSURE\x10\x04\x12,\n(KUBE_NODE_CONDITION_TYPE_NETWORK_UNAVAIL\x10\x05*\xa4\x01\n\rNodeAdmission\x12\x1e\n\x1aNODE_ADMISSION_UNSPECIFIED\x10\x00\x12 \n\x1cNODE_ADMISSION_NOT_CLUSTERED\x10\x01\x12\x1a\n\x16NODE_ADMISSION_LEAVING\x10\x02\x12\x1a\n\x16NODE_ADMISSION_JOINING\x10\x03\x12\x19\n\x15NODE_ADMISSION_JOINED\x10\x04*\x99\x03\n\x13StorageHealthStatus\x12%\n!STORAGE_HEALTH_STATUS_UNSPECIFIED\x10\x00\x12!\n\x1dSTORAGE_HEALTH_STATUS_HEALTHY\x10\x01\x12\x42\n>STORAGE_HEALTH_STATUS_DEGRADED_2_REPLICA_AVAILABLE_REPLICATING\x10\x02\x12\x46\nBSTORAGE_HEALTH_STATUS_DEGRADED_2_REPLICA_AVAILABLE_NOT_REPLICATING\x10\x03\x12\x42\n>STORAGE_HEALTH_STATUS_DEGRADED_1_REPLICA_AVAILABLE_REPLICATING\x10\x04\x12\x46\nBSTORAGE_HEALTH_STATUS_DEGRADED_1_REPLICA_AVAILABLE_NOT_REPLICATING\x10\x05\x12 \n\x1cSTORAGE_HEALTH_STATUS_FAILED\x10\x06*\x8e\x02\n\x12StorageVolumeState\x12$\n STORAGE_VOLUME_STATE_UNSPECIFIED\x10\x00\x12!\n\x1dSTORAGE_VOLUME_STATE_CREATING\x10\x01\x12!\n\x1dSTORAGE_VOLUME_STATE_ATTACHED\x10\x02\x12!\n\x1dSTORAGE_VOLUME_STATE_DETACHED\x10\x03\x12\"\n\x1eSTORAGE_VOLUME_STATE_ATTACHING\x10\x04\x12\"\n\x1eSTORAGE_VOLUME_STATE_DETACHING\x10\x05\x12!\n\x1dSTORAGE_VOLUME_STATE_DELETING\x10\x06*\xba\x01\n\x17StorageVolumeRobustness\x12)\n%STORAGE_VOLUME_ROBUSTNESS_UNSPECIFIED\x10\x00\x12%\n!STORAGE_VOLUME_ROBUSTNESS_HEALTHY\x10\x01\x12&\n\"STORAGE_VOLUME_ROBUSTNESS_DEGRADED\x10\x02\x12%\n!STORAGE_VOLUME_ROBUSTNESS_FAULTED\x10\x03*\x86\x02\n\x16StorageVolumePVCStatus\x12)\n%STORAGE_VOLUME_PVC_STATUS_UNSPECIFIED\x10\x00\x12#\n\x1fSTORAGE_VOLUME_PVC_STATUS_BOUND\x10\x01\x12%\n!STORAGE_VOLUME_PVC_STATUS_PENDING\x10\x02\x12\'\n#STORAGE_VOLUME_PVC_STATUS_AVAILABLE\x10\x03\x12&\n\"STORAGE_VOLUME_PVC_STATUS_RELEASED\x10\x04\x12$\n STORAGE_VOLUME_PVC_STATUS_FAILED\x10\x05*\xd0\x02\n\x1aStorageVolumeReplicaStatus\x12-\n)STORAGE_VOLUME_REPLICA_STATUS_UNSPECIFIED\x10\x00\x12,\n(STORAGE_VOLUME_REPLICA_STATUS_REBUILDING\x10\x01\x12(\n$STORAGE_VOLUME_REPLICA_STATUS_ONLINE\x10\x02\x12(\n$STORAGE_VOLUME_REPLICA_STATUS_FAILED\x10\x03\x12)\n%STORAGE_VOLUME_REPLICA_STATUS_OFFLINE\x10\x04\x12*\n&STORAGE_VOLUME_REPLICA_STATUS_STARTING\x10\x05\x12*\n&STORAGE_VOLUME_REPLICA_STATUS_STOPPING\x10\x06*\xac\x02\n\rKubePodStatus\x12\x1f\n\x1bKUBE_POD_STATUS_UNSPECIFIED\x10\x00\x12\x1b\n\x17KUBE_POD_STATUS_PENDING\x10\x01\x12\x1b\n\x17KUBE_POD_STATUS_RUNNING\x10\x02\x12\x1d\n\x19KUBE_POD_STATUS_SUCCEEDED\x10\x03\x12&\n\"KUBE_POD_STATUS_CONTAINER_CREATING\x10\x04\x12%\n!KUBE_POD_STATUS_CRASHLOOP_BACKOFF\x10\x05\x12\x19\n\x15KUBE_POD_STATUS_ERROR\x10\x06\x12\x1b\n\x17KUBE_POD_STATUS_EVICTED\x10\x07\x12\x1a\n\x16KUBE_POD_STATUS_FAILED\x10\x08*\x83\x01\n\rServiceStatus\x12\x1e\n\x1aSERVICE_STATUS_UNSPECIFIED\x10\x00\x12\x19\n\x15SERVICE_STATUS_FAILED\x10\x01\x12\x1b\n\x17SERVICE_STATUS_DEGRADED\x10\x02\x12\x1a\n\x16SERVICE_STATUS_HEALTHY\x10\x03*\x86\x02\n\x14KubeCompUpdateStatus\x12\'\n#KUBE_COMP_UPDATE_STATUS_UNSPECIFIED\x10\x00\x12$\n KUBE_COMP_UPDATE_STATUS_DOWNLOAD\x10\x01\x12+\n\'KUBE_COMP_UPDATE_STATUS_DOWNLOAD_FAILED\x10\x02\x12\'\n#KUBE_COMP_UPDATE_STATUS_IN_PROGRESS\x10\x03\x12\"\n\x1eKUBE_COMP_UPDATE_STATUS_FAILED\x10\x04\x12%\n!KUBE_COMP_UPDATE_STATUS_COMPLETED\x10\x05*\xab\x01\n\x08KubeComp\x12\x19\n\x15KUBE_COMP_UNSPECIFIED\x10\x00\x12\x18\n\x14KUBE_COMP_CONTAINERD\x10\x01\x12\x11\n\rKUBE_COMP_K3S\x10\x02\x12\x14\n\x10KUBE_COMP_MULTUS\x10\x03\x12\x16\n\x12KUBE_COMP_KUBEVIRT\x10\x04\x12\x11\n\rKUBE_COMP_CDI\x10\x05\x12\x16\n\x12KUBE_COMP_LONGHORN\x10\x06*\x81\x02\n\rKubeVMIStatus\x12\x1f\n\x1bKUBE_VMI_STATUS_UNSPECIFIED\x10\x00\x12\x1b\n\x17KUBE_VMI_STATUS_PENDING\x10\x01\x12\x1e\n\x1aKUBE_VMI_STATUS_SCHEDULING\x10\x02\x12\x1d\n\x19KUBE_VMI_STATUS_SCHEDULED\x10\x03\x12\x1b\n\x17KUBE_VMI_STATUS_RUNNING\x10\x04\x12\x1d\n\x19KUBE_VMI_STATUS_SUCCEEDED\x10\x05\x12\x1a\n\x16KUBE_VMI_STATUS_FAILED\x10\x06\x12\x1b\n\x17KUBE_VMI_STATUS_UNKNOWN\x10\x07\x42\x39\n\x13org.lfedge.eve.infoZ\"github.com/lf-edge/eve-api/go/infob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'info.edge_node_cluster_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023org.lfedge.eve.infoZ\"github.com/lf-edge/eve-api/go/info'
  _globals['_KUBENODECONDITIONTYPE']._serialized_start=2031
  _globals['_KUBENODECONDITIONTYPE']._serialized_end=2311
  _globals['_NODEADMISSION']._serialized_start=2314
  _globals['_NODEADMISSION']._serialized_end=2478
  _globals['_STORAGEHEALTHSTATUS']._serialized_start=2481
  _globals['_STORAGEHEALTHSTATUS']._serialized_end=2890
  _globals['_STORAGEVOLUMESTATE']._serialized_start=2893
  _globals['_STORAGEVOLUMESTATE']._serialized_end=3163
  _globals['_STORAGEVOLUMEROBUSTNESS']._serialized_start=3166
  _globals['_STORAGEVOLUMEROBUSTNESS']._serialized_end=3352
  _globals['_STORAGEVOLUMEPVCSTATUS']._serialized_start=3355
  _globals['_STORAGEVOLUMEPVCSTATUS']._serialized_end=3617
  _globals['_STORAGEVOLUMEREPLICASTATUS']._serialized_start=3620
  _globals['_STORAGEVOLUMEREPLICASTATUS']._serialized_end=3956
  _globals['_KUBEPODSTATUS']._serialized_start=3959
  _globals['_KUBEPODSTATUS']._serialized_end=4259
  _globals['_SERVICESTATUS']._serialized_start=4262
  _globals['_SERVICESTATUS']._serialized_end=4393
  _globals['_KUBECOMPUPDATESTATUS']._serialized_start=4396
  _globals['_KUBECOMPUPDATESTATUS']._serialized_end=4658
  _globals['_KUBECOMP']._serialized_start=4661
  _globals['_KUBECOMP']._serialized_end=4832
  _globals['_KUBEVMISTATUS']._serialized_start=4835
  _globals['_KUBEVMISTATUS']._serialized_end=5092
  _globals['_KUBENODECONDITION']._serialized_start=87
  _globals['_KUBENODECONDITION']._serialized_end=235
  _globals['_KUBENODEINFO']._serialized_start=238
  _globals['_KUBENODEINFO']._serialized_end=552
  _globals['_KUBEPODNAMESPACEINFO']._serialized_start=555
  _globals['_KUBEPODNAMESPACEINFO']._serialized_end=719
  _globals['_KUBEEVEAPPPODINFO']._serialized_start=722
  _globals['_KUBEEVEAPPPODINFO']._serialized_end=980
  _globals['_KUBEVOLUMEREPLICAINFO']._serialized_start=983
  _globals['_KUBEVOLUMEREPLICAINFO']._serialized_end=1142
  _globals['_KUBEVOLUMEINFO']._serialized_start=1145
  _globals['_KUBEVOLUMEINFO']._serialized_end=1621
  _globals['_KUBESTORAGEINFO']._serialized_start=1624
  _globals['_KUBESTORAGEINFO']._serialized_end=1800
  _globals['_KUBEVMIINFO']._serialized_start=1803
  _globals['_KUBEVMIINFO']._serialized_end=2028
# @@protoc_insertion_point(module_scope)