CURRENT_YEAR=$(date +"%Y")
NEXT_YEAR=$((CURRENT_YEAR+1))

_sn_a6x="cfg/base.yaml,cfg/sn_a6x.base.yaml"

_configurations=(
  2 "${_sn_a6x},cfg/template_planner_daily_with_cal.yaml,cfg/sn_a6x.planner.daily-with-cal.yaml"
)

_configurations_len=${#_configurations[@]}

for _year in $NEXT_YEAR; do
  for _idx in $(seq 0 2 $((_configurations_len-1))); do
    _passes=${_configurations[_idx]}
    _cfg=${_configurations[_idx+1]}

    PLANNER_YEAR=${_year} PASSES=${_passes} CFG="${_cfg}" ./single.sh
  done
done
