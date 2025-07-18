# Key Pair Value Generator

This app will generate rsa 2048 key value pair to be used in RS256 crypto processes, with `n` and `e` values to regenerate the public key.

An example of response can be found here:
```json
{
  "private_key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeTYrSkR1b2RUT24zS2VONHd2M3JqYVJXWVRoWnREbm5CWXBwOHlDeHZNcm5RNE9wCndFVzBmVDlpbEs5RnVkOGVmT0hDeTZVcWR3TFVsemtLb1ZGT0NxeldtODZLWWpFR0U1eHFacmlUM2U4YUN3VlkKTDRhZnp0c3dlNFA0UlNMOWpsTEpKd1c2VGhHTmdhT0dSKy9JQTdxU2FkSXpteEVtNElWdmd5NktBSzk0QXpMaQo2SU5DOXhjbEFkV2ZsUDE1VFNKRUtpUkgwRldnRlo4YjhpWmFhMkRXT08vbVdjN0VVSk5IZTh6YkpTcFRESktlCkJlbVlacnRnN2hyZUNXYWRjMlZRYWZaMmp1RUhTbUZ0TEd1Y3RjNkQyUVdPcUU2QThkMW1NS3Z1L0RTUEFRWEIKT2JKaTVmaVo5OFZ1blMwQ2d5N2phSU04eUhzb2V1OGVsSDhVVlFJREFRQUJBb0lCQUJqbzRlaWhNMWg0SzllRApiNmlIellVTmY0T0tsTmN6SzVza1hIcGZoR0pwS0JPUmZiTG1IaGJMMG41ZmhHQ2hDNXFoMjNwSjVveHJ4ckdyCkhWbmNLOHRTRnVReXcvNjlMY0tST0ZGUlZSZ1RzOTk4S3dHVG1SQ3ByVjlnVDM1WkNLR2RZenlkVG9zRW40WFIKWGZjeDQ3S250YkhicmZSU3RTcWVkdm1YNC9NRm9BR29zME04WGVNaUFCcngvTmpXY3RvWnFDZ3BBQitjc0JlNQpMdWFIT1JCWHd2NFJVU2c1Y2hhck1YNXpuSU50Wjh6YTc5S1J6Zy9Qay9scHZuY3dCS3d0N3d1QXNKSllzYWFQCm5OckphaGxDY1VKRkptMEpjaDhPdDBUbGtXU3NHY1YwV2VSUEE2NkRiaW5nQ2tTbzhpOXpLeGNLRXdLRUcvcVcKWDN0SVdYTUNnWUVBNm1xRDBzM2hjbmlVMW92V0U3WHdTQVQ1MDVqa0JkK3M0clRjb0ZTaWVvQUNMZG9GakI3QgpUUFZiM2lkSkt0VUt0NTloQU1hWTFnVmJaRWhpcXV2SzVoNkI5TE5ETjZSSWp0RDVoRWZNM0tsTHcxbzBoY1BGCkhDVkVMMXBNdVFjWVRPd0J0OFBsdWZmdTdjdGR6UmRXZnZUZFZja1dkVkFWcnJ5bEdPbFVWUWNDZ1lFQTNuQ3EKaE9oOTJGNGFkWnlGYXJONE5EckEwaW0zSEYwTkJaMW9DejhIeHNrRjVPdUNSQ25oa3k4SWlkb1RwRCt2N0p5VwpHdHF1bGlJa3lxSTV4dEU3cWY3TjZINmkrWkpoOVZ3ZnJ3WmZkL0U4d2lwYzJzT25Cai9FSXpOcnJLUkhkRnJvCnIwRVpwSTJWYmRqays3ZVlSdU5Ec0l4MmUzYmV2U2RERFREbk1NTUNnWUIwZXVyMzg4K0Rwd0VtUHFQb2RXNXAKQ3Zmc3Vic01aQ0d2SVBuRVBXbmkvdnFXT0JDcm1KaXFtZnpGUGJZd3IxMjg4bGdzSDRMUVRpY2toSWRxc1BISQpPUDZRVGdjbmZkMkVBYmtLanZidjZydTMzWG5kd2ZLTzBzRzMyZUhueXV1N1JVWnhQc2xIQ0RqVU5rcHMzNXplClJ1UTRmVXhJakx2SmNQbUVrR1NVUHdLQmdRQ0wzVnZhNWpUNjczRERzdjlROERnTTMwU080UnY2QWFoR2Q0c00KUnoya2lKSGtOeEZadXR2VHFDYytScGxERmFpUXBvVGJyZFZKSVg3d3lzQjVldWFCSlc1aldmMHY0Qk1vTEJ6ZgorS25leVJqakhhLzAva0R2ZFBqbkNWWVMzOWp1R0p0bVZ4REFueVh1UkFyM0c2SHNNTXBISUlyL1p1YjNHTkRZClluOTc4UUtCZ0JOVytNM2RDaHJjN05PVEVXQS8vZ0ViUDQ5M3FkdDJMTjF5MGYzMll3TllBcEVLdWRscFhEMXoKaEVCazRSSHVyUnIvb2RUekVaLzA4Q1Q1TGZycktsd2NDcjg4MWNMMzRMQ2ppVjdOc0w4V0NDNm85cVFsT2RqWApiMXczTXVaZmlTbUpBY2Y0Ynl2VlVIN3FvUTg5amhiVjFOVm5VQjdmRFI1bUZjYVdpRXJpCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
  "public_key": "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF5NitKRHVvZFRPbjNLZU40d3YzcgpqYVJXWVRoWnREbm5CWXBwOHlDeHZNcm5RNE9wd0VXMGZUOWlsSzlGdWQ4ZWZPSEN5NlVxZHdMVWx6a0tvVkZPCkNxeldtODZLWWpFR0U1eHFacmlUM2U4YUN3VllMNGFmenRzd2U0UDRSU0w5amxMSkp3VzZUaEdOZ2FPR1IrL0kKQTdxU2FkSXpteEVtNElWdmd5NktBSzk0QXpMaTZJTkM5eGNsQWRXZmxQMTVUU0pFS2lSSDBGV2dGWjhiOGlaYQphMkRXT08vbVdjN0VVSk5IZTh6YkpTcFRESktlQmVtWVpydGc3aHJlQ1dhZGMyVlFhZloyanVFSFNtRnRMR3VjCnRjNkQyUVdPcUU2QThkMW1NS3Z1L0RTUEFRWEJPYkppNWZpWjk4VnVuUzBDZ3k3amFJTTh5SHNvZXU4ZWxIOFUKVlFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
  "n": "25712935632000259286458586683415614235022584134216574613743727798489992446041082241484709773163162103537420942676102127139194703614940891397389768140574737467994293099179867299435282700031882706238276936506207653443670880311802534285835724417212317653789862134339593620896432332545830347225766360816448749661089970200473553757444377556059474186603289655651913244073435022803923590227126782551180878908024308863376542782103183020160207458056186899522386571960278880814227901758043713255631255346076055112058698971448378031140882368816445571362831107618101683819505740578419277163919520903131415466730675540741008069717",
  "e": 65537,
  "alg": "RS256"
}
```

Private and Public keys are encoded is a string in base64 format 

## How to build
Building process of this repository happens calling `.\build.ps1` that build the final executable to be included or called where it needs

## How to run locally
Run those in your console
- `docker build . -t key-pair`
- `docker run -it --rm key-pair`

And in console you will get a json with the generated key value pairs