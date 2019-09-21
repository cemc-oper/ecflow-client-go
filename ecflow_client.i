%module ecflow_client
%{
#include "ecflow_client.h"
%}

%include "std_vector.i"
%include "std_string.i"

%include "ecflow_client.h"

namespace std {
  %template(NodeStatusRecordVector) vector<EcflowUtil::NodeStatusRecord>;
};