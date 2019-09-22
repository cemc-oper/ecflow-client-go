%module ecflow_client
%{
#include "ecflow_util.h"
%}

%include "std_vector.i"
%include "std_string.i"

%include "ecflow_util.h"

namespace std {
  %template(NodeStatusRecordVector) vector<EcflowUtil::NodeStatusRecord>;
};